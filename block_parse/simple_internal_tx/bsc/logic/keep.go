package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"kryptos_rai/block_parse/simple_internal_tx/bsc/config"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	"kryptos_rai/use/pool"
	"math/big"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

func GetInternalTx() {

	bodyMap := make(map[string]interface{})
	bodyMap["jsonrpc"] = "2.0"
	bodyMap["method"] = "eth_getLogs"
	bodyMap["id"] = 74

	fromBlock := config.GetStartBlock()
	fromHex := "0x" + strconv.FormatInt(int64(fromBlock), 16)
	toBlock := int64(0)
	toBlockHex := ""

	index := 0
	tmpNodeURL := config.GetBSCNodeURL()
	indexList := Perm(len(tmpNodeURL))
	nodeURL := []string{}
	for _, x := range indexList {
		nodeURL = append(nodeURL, tmpNodeURL[x])
	}
	nodeLen := len(nodeURL)

	for {
		//上限不超过5000 但是测试发现就连1000的跨度都不行 100是可以的

		if toBlock >= 0 {
			_, latestBlockNumber := getLatestBlock()
			if latestBlockNumber-fromBlock >= 99 {
				toBlock = fromBlock + 99
				toBlockHex = "0x" + strconv.FormatInt(toBlock, 16)
			} else {
				toBlock = -1
				toBlockHex = "latest"
			}
		}

		paramMap := map[string]interface{}{"topics": []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}}
		paramMap["fromBlock"] = fromHex
		paramMap["toBlock"] = toBlockHex

		paramList := []map[string]interface{}{}
		paramList = append(paramList, paramMap)
		bodyMap["params"] = paramList
		rawBody, _ := json.Marshal(bodyMap)

		for {
			client := pool.Pool.Get().(http.Client)

			req, _ := http.NewRequest("POST", nodeURL[index], bytes.NewBuffer(rawBody))
			defer req.Body.Close()
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("content-encoding", "gzip")

			r, err := client.Do(req)
			if err != nil || r == nil {
				logger.Error("GetInternalTx::failed.try another node %+v", err)
				index = (index + 1) % nodeLen
				pool.Pool.Put(client)
				continue
			}

			if r != nil && r.Body != nil {
				defer r.Body.Close()
			} else {
				logger.Error("GetInternalTx:: result is nil")
				index = (index + 1) % nodeLen
				pool.Pool.Put(client)
				continue
			}

			retData := new(NodeRet)
			retData.NodeRetInit()
			retBytes, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal([]byte(retBytes), &retData)

			code, ok1 := retData.Error["code"]
			msg, ok2 := retData.Error["message"]
			if ok1 || ok2 {
				logger.Error("GetInternalTx:: failed code:%+v,message:%+v,url:%+v", code, msg, nodeURL[index])
				index = (index + 1) % nodeLen
				pool.Pool.Put(client)
				continue
			}

			retLen := len(retData.Result)
			logger.Notic("GetInternalTx:: %+v,from:%+v,to:%+v", retLen, fromBlock, toBlock)
			if retLen > 0 {
				insertList := InTxRecordList{}
				updateTime := time.Now().Unix()
				for _, v := range retData.Result {
					if len(v.Topics) < 3 {
						continue
					}
					item := InTxRecord{}
					item.Address = v.Address
					num, _ := strconv.ParseInt(v.BlockNumber, 0, 64)
					item.BlockNumber = num

					item.From = "0x" + v.Topics[1][26:]
					item.To = "0x" + v.Topics[2][26:]
					item.TxHash = v.TxHash
					item.UpdataTime = updateTime
					value, _ := new(big.Int).SetString(v.Data, 0)
					item.Value = value.String()
					insertList = append(insertList, item)
				}
				sort.Sort(insertList)

				//找最新的那一条
				tmpFromBlock := fromBlock

				new := retData.Result[len(insertList)-1].BlockNumber
				newNumber, err := strconv.ParseInt(new, 0, 64)
				if err != nil {
					logger.Error("GetInternalTx::ParseInt failed :%+v", err)
					index = (index + 1) % nodeLen
					pool.Pool.Put(client)
					continue
				}

				fromBlock = newNumber + 1
				fromHex = "0x" + strconv.FormatInt(fromBlock, 16)

				ids, err := mongo.InternalTxClient.InsertMany(context.Background(), insertList)
				if err != nil {
					logger.Error("GetInternalTx::insert many failed:%+v,ids len is :%+v,insert lenis:%+v", err, len(ids.InsertedIDs), len(insertList))
				}
				logger.Notic("GetInternalTx::inser %+v record success,fromBlock:%+v,toBlock:%+v,toBlockHex:%+v,latest insert is :%+v", len(ids.InsertedIDs), tmpFromBlock, toBlock, toBlockHex, newNumber)

				//为了测试只保留10万个区块数据
				u := bson.M{}
				u["block_number"] = bson.M{"$lte": newNumber - 100000}
				err = mongo.InternalTxClient.Remove(context.Background(), u)
				if err != nil {
					logger.Warn("GetInternalTx::remove failed:%+v,data is :%+v", err, u)
				}

			} else {
				if toBlock > 0 {
					fromBlock = toBlock + 1
					fromHex = "0x" + strconv.FormatInt(fromBlock, 16)
				} else {
					//连续几百个区块没交易 from一直保持在原地
					_, latestBlockNumber := getLatestBlock()
					if fromBlock < latestBlockNumber-100 {
						fromBlock = latestBlockNumber - 100
						fromHex = "0x" + strconv.FormatInt(fromBlock, 16)
					}
				}
			}

			index = (index + 1) % nodeLen
			pool.Pool.Put(client)
			break
		}
		time.Sleep(time.Second * time.Duration(config.GetCheckInterval()))
	}
}

func getLatestBlock() (hexstring string, blockNumber int64) {

	index := 0
	tmpNodeURL := config.GetBSCNodeURL()
	indexList := Perm(len(tmpNodeURL))
	nodeURL := []string{}
	for _, x := range indexList {
		nodeURL = append(nodeURL, tmpNodeURL[x])
	}
	nodeLen := len(nodeURL)

	client := pool.Pool.Get().(http.Client)
	for {

		bodyMap := make(map[string]interface{})
		bodyMap["jsonrpc"] = "2.0"
		bodyMap["method"] = "eth_blockNumber"
		bodyMap["params"] = []interface{}{}
		bodyMap["id"] = 83
		rawBody, _ := json.Marshal(bodyMap)
		req, _ := http.NewRequest("POST", nodeURL[index], bytes.NewBuffer(rawBody))
		if req != nil && req.Body != nil {
			defer req.Body.Close()
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil || resp == nil {
			logger.Error("getBSCLatestBlock::Get failed %+v,resp:%+v,url:%+v", err, resp, nodeURL[index])
			pool.Pool.Put(client)
			index = (index + 1) % nodeLen
			continue
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		//{"jsonrpc":"2.0","id":83,"result":"0x81734b"}
		retData := new(GetBlockRet)
		retData.GetBlockRetInit()
		retBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(retBytes), retData)
		logger.Info("getBSCLatestBlock:: result:%+v", retData)

		if retData.Result != "" {
			num, err := strconv.ParseInt(retData.Result, 0, 64)
			if err != nil {
				logger.Error("getBSCLatestBlock::ParseInt failed %+v,url:%+v", err, nodeURL[index])
				pool.Pool.Put(client)
				index = (index + 1) % nodeLen
				continue
			}
			pool.Pool.Put(client)
			return retData.Result, num

		} else {
			logger.Error("getBSCLatestBlock::result is nil,url:%+v", nodeURL[index])
			pool.Pool.Put(client)
			index = (index + 1) % nodeLen
			continue
		}
	}
}
