package logic

import "sync"

var curBlock int64

type NodeRet struct {
	JSONRPC string `json:jsonrpc`
	ID      int    `json:id`
	Result  string `json:result`
}

func (s *NodeRet) NodeRetInit() {
	s.JSONRPC = ""
	s.ID = 0
	s.Result = ""
}

type TxData struct {
	BlockNumber int64  `bson:block_number`
	Hash        string `bson:"hash"`
	From        string `bson:"from"`
	To          string `bson:to`
	UpdataTime  int64  `bson:"update_time"`
}

type RPCTransaction struct {
	From string `json:"from"`
	Hash string `json:"hash"`
	To   string `json:"to"`
}

type BlockResult struct {
	Number       string           `json:"number"`
	Transactions []RPCTransaction `json:"transactions"`
}

type Block struct {
	Result BlockResult `json:"result"`
}

func (b *Block) InitBlock() {
	result := BlockResult{}
	result.Number = ""
	result.Transactions = []RPCTransaction{}
	b.Result = result
}

type TxDataList []TxData

func (a TxDataList) Len() int           { return len(a) }
func (a TxDataList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TxDataList) Less(i, j int) bool { return a[i].BlockNumber < a[j].BlockNumber }

type concurrencyList struct {
	Mu     sync.Mutex
	TxList TxDataList
}
