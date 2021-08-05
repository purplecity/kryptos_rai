package app

type SubscribeInfoReq struct {
	PoolAddress     string `json:"pool_address"`
	MasterChelfAddr string `json:"master_chelf_addr"`
	UserAddress     string `json:"user_address"`
	Email           string `json:"email"`
	WatchFlashLoan  string `json:"watch_fl"`
	TVLInc          int64  `json:"tvl_inc"`
	TVLDec          int64  `json:"tvl_dec"`
	Method          string `json:"method"` //操作方法
}

type LogoReq struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	LogoType     string `json:"type"`
	TokenAddress string `json:"token_address"`
	Chain        string `json:"chain"`
	ID           string `json:"_id"`
	Method       string `json:"method"` //操作方法
	Page         int64  `json:"page"`   //查询页码

}

type OperatePoolInfoReq struct {
	Platform           string   `json:"platform"`              //平台名称
	Chain              string   `json:"chain"`                 //链名称
	InputTokenList     []string `json:"input_token_list"`      //输入token列表
	CertificateToken   string   `json:"certificate_token"`     //凭证token名称
	PoolAddress        string   `json:"pool_address"`          //池子地址
	FundFlowList       []string `json:"fund_flow_list"`        //资金流向列表
	PoolProperty       string   `json:"pool_property"`         //池子性质
	IsWatching         string   `json:"is_watching"`           //是否监控
	WatchDescribe      []string `json:"watch_describe"`        //监控描述
	Method             string   `json:"method"`                //操作方法
	Page               int64    `json:"page"`                  //查询页码
	InputTokenNameList []string `json:"input_token_name_list"` //输入token名列表主要为了查询
	QueryTokenName     string   `json:"query_token_name"`      //输入token名列表主要为了查询
	RewardTokenList    []string `json:"reward_token_list"`     //奖励token列表
}

type Response struct {
	Code int    // 0 success  others error
	Msg  string // success  errorMsg
	Data interface{}
}

func (r *Response) ResponseInit() {
	r.Code = SuccessCode
	r.Msg = SuccessMsg
	r.Data = map[string]interface{}{}
}

type AddressAnalysisReq struct {
	Address        string `json:"address"`
	Chain          string `json:"chain"`
	ConnectAddress string `json:"connect_address"`
}

type subData struct {
	HasBalance      bool
	MasterChelfAddr string
}
