package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"kryptos_rai/block_parse/simple_tx/bsc/config"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	"kryptos_rai/use/pool"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
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

func KeepBSCLatest() {
	for {
		latest := getLatestBlock()
		if curBlock >= latest {
			time.Sleep(time.Second * time.Duration(config.GetCheckInterval()))
			continue
		}

		if latest-curBlock >= 100 {
			latest = curBlock + 100
		}

		start := curBlock
		end := latest
		count := (end - start) / 5
		mod := (end - start) % 5
		wg := new(sync.WaitGroup)
		txList := new(concurrencyList)

		//5个一组
		for i := int64(0); i < count; i++ {
			wg.Add(1)
			t_start := i*5 + start + 1
			t_end := t_start + 4
			go upToLatest(t_start, t_end, wg, txList)

		}
		start += 5 * count

		if mod != 0 {
			wg.Add(1)
			upToLatest(start+1, latest, wg, txList)
		}

		wg.Wait()

		sort.Sort(txList.TxList)
		if len(txList.TxList) > 0 {
			result, err := mongo.TxMongoClient.InsertMany(context.Background(), txList.TxList)
			if err != nil || result == nil || len(result.InsertedIDs) != len(txList.TxList) {
				logger.Fatal("InsertBlockTx:: InsertMany failed:%+v,len:%+v,result:%+v", err, txList.TxList, result)
				return
			}
		}

		//为了测试 只保留100000个区块的数据
		// u := bson.M{}
		// u["block_number"] = bson.M{"$lt": latest - 100000}
		// err := mongo.TxMongoClient.Remove(context.Background(), u)
		// if err != nil {
		// 	logger.Warn("Remove Block Tx::remove failed:%+v,data is :%+v", err, u)
		// }

		logger.Info("Block interval finished::curBlock:%+v,latest:%+v, total tx is :%+v", curBlock, latest, len(txList.TxList))
		curBlock = latest
		time.Sleep(time.Second * time.Duration(config.GetCheckInterval()))
	}

}

func upToLatest(start, end int64, wg *sync.WaitGroup, txList *concurrencyList) {
	for t_start := start; t_start <= end; t_start++ {
		s := "0x" + strconv.FormatInt(t_start, 16)
		err := parseBlock(s, txList)
		if err != nil {
			logger.Error("upToLatest::ParseBlock failed,check it %+v,hex:%+v", t_start, s)
			wg.Done()
			return
		}
	}
	wg.Done()
}

func parseBlock(blockHexBlock string, txList *concurrencyList) error {
	index := 0
	tmpNodeURL := config.GetBSCNodeURL()
	indexList := Perm(len(tmpNodeURL))
	nodeURL := []string{}
	for _, x := range indexList {
		nodeURL = append(nodeURL, tmpNodeURL[x])
	}
	nodeLen := len(nodeURL)

	for {
		client := pool.Pool.Get().(http.Client)
		/*
				{
				"jsonrpc":"2.0",
				"method":"eth_getBlockByNumber",
				"params":[
					"0x1b4",
					true
				],
				"id":1
			}
		*/
		bodyMap := make(map[string]interface{})
		bodyMap["jsonrpc"] = "2.0"
		bodyMap["method"] = "eth_getBlockByNumber"
		bodyMap["id"] = 1
		bodyMap["params"] = []interface{}{blockHexBlock, true}
		rawBody, _ := json.Marshal(bodyMap)
		req, _ := http.NewRequest("POST", nodeURL[index], bytes.NewBuffer(rawBody))
		defer req.Body.Close()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("content-encoding", "gzip")

		r, err := client.Do(req)
		if err != nil || r == nil {
			logger.Error("ParseBlock::request failed. err:%+v,url:%+v", err, nodeURL[index])
			index = (index + 1) % nodeLen
			pool.Pool.Put(client)
			continue
		}

		if r != nil && r.Body != nil {
			defer r.Body.Close()
		} else {
			logger.Error("ParseBlock::body is nil. url:%+v", nodeURL[index])
			index = (index + 1) % nodeLen
			pool.Pool.Put(client)
			continue
		}

		retData := new(Block)
		retData.InitBlock()
		retBytes, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal([]byte(retBytes), &retData)

		num := retData.Result.Number
		num10, err := strconv.ParseInt(num, 0, 64)
		if err != nil {
			logger.Error("ParseBlock: parseint failed:%+v,ts:%+v,url:%+v", err, num, nodeURL[index])
			index = (index + 1) % nodeLen
			pool.Pool.Put(client)
			continue
		}

		data := []TxData{}
		timeNow := time.Now().Unix()
		for _, e := range retData.Result.Transactions {
			txData := TxData{
				BlockNumber: num10,
				Hash:        e.Hash,
				From:        e.From,
				To:          e.To,
				UpdataTime:  timeNow,
			}
			data = append(data, txData)
		}
		txList.Mu.Lock()
		txList.TxList = append(txList.TxList, data...)
		txList.Mu.Unlock()
		pool.Pool.Put(client)
		return nil
	}
}

func getLatestBlock() int64 {

	index := 0
	tmpNodeURL := config.GetBSCNodeURL()
	indexList := Perm(len(tmpNodeURL))
	nodeURL := []string{}
	for _, x := range indexList {
		nodeURL = append(nodeURL, tmpNodeURL[x])
	}
	nodeLen := len(nodeURL)

	for {

		client := pool.Pool.Get().(http.Client)

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
		retData := new(NodeRet)
		retData.NodeRetInit()
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
			return num

		} else {
			logger.Error("getBSCLatestBlock::result is nil,url:%+v", nodeURL[index])
			pool.Pool.Put(client)
			index = (index + 1) % nodeLen
			continue
		}
	}
}
