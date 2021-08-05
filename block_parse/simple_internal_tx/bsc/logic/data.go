package logic

/*
{
    "jsonrpc": "2.0",
    "id": 74,
    "error": {
        "code": -32602,
        "message": "invalid argument 0: hex string of odd length"
    }
}


        {
            "address": "0xf9cec8d50f6c8ad3fb6dccec577e05aa32b224fe",
            "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x000000000000000000000000d488e1959360dd9ab02a8a9488ccc5a094290cf5",
                "0x0000000000000000000000006045931e511ef7e53a4a817f971e0ca28c758809"
            ],
            "data": "0x00000000000000000000000000000000000000000000000000000008ae8b203f",
            "blockNumber": "0x8f01eb",
            "transactionHash": "0x701792f89aac0d3de22b73b08e31902e2ea222f0d78f550557fb56b092d349e5",
            "transactionIndex": "0x0",
            "blockHash": "0x4bc87b4e57d1945b40939ffa487059d09a24169f56d4697dae77143edd543f22",
            "logIndex": "0x0",
            "removed": false
        },


*/
type NodeRet struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Result  []InTx                 `json:"result"`
	Error   map[string]interface{} `json:"error"`
}

func (s *NodeRet) NodeRetInit() {
	s.JSONRPC = ""
	s.ID = 0
	s.Result = []InTx{}
	s.Error = map[string]interface{}{}
}

type InTx struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber string   `json:"blockNumber"`
	TxHash      string   `json:"transactionHash"`
}

type InTxRecord struct {
	Address     string `bson:"address"`
	From        string `bson:"from"`
	To          string `bson:"to"`
	Value       string `bson:"value"`
	BlockNumber int64  `bson:"block_number"`
	TxHash      string `bson:"transaction_hash"`
	UpdataTime  int64  `bson:"update_time"`
}

type InTxRecordList []InTxRecord

func (a InTxRecordList) Len() int           { return len(a) }
func (a InTxRecordList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InTxRecordList) Less(i, j int) bool { return a[i].BlockNumber < a[j].BlockNumber }

type GetBlockRet struct {
	JSONRPC string `json:jsonrpc`
	ID      int    `json:id`
	Result  string `json:result`
}

func (s *GetBlockRet) GetBlockRetInit() {
	s.JSONRPC = ""
	s.ID = 0
	s.Result = ""
}
