package watch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pool"
	"net/http"
	"strconv"
	"time"
)

type nodeRet struct {
	JSONRPC string `json:jsonrpc`
	ID      int    `json:id`
	Result  string `json:result`
}

func (s *nodeRet) nodeRetInit() {
	s.JSONRPC = ""
	s.ID = 0
	s.Result = ""
}

func GetBSCLatestBlock() int64 {

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
		retData := new(nodeRet)
		retData.nodeRetInit()
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
			if num > bscBlockHeight {
				bscBlockHeight = num
				for _, p := range pool_info.BscPoolList {
					p.WorkChan <- num
				}
			}
			time.Sleep(time.Second * time.Duration(config.GetBSCCheckInterval()))
			return num

		} else {
			logger.Error("getBSCLatestBlock::result is nil,url:%+v", nodeURL[index])
			pool.Pool.Put(client)
			index = (index + 1) % nodeLen
			continue
		}
	}
}
