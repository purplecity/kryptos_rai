package app

type SubscribeInfo struct {
	//field.DefaultField `bson:",inline"`
	PoolAddress     string `bson:"pool_address"`      //池子地址 小写
	MasterChelfAddr string `bson:"master_chelf_addr"` //masterchelf类型地址 小写
	UserAddress     string `bson:"user_address"`      //用户链接地址 没有账户体系 把metamask链接地址当作账户 小写
	Email           string `bson:"email"`
	WatchFlashLoan  string `bson:"watch_fl"`
	TVLInc          int64  `bson:"tvl_inc"`
	TVLDec          int64  `bson:"tvl_dec"`
	UpdateTime      int64  `bson:"update_time"` //更新时间
}

type Logo struct {
	Name         string `bson:"name"` // 小写
	URL          string `bson:"url"`
	LogoType     string `bson:"type"`          //项目 币种
	TokenAddress string `bson:"token_address"` //小写
	Chain        string `bson:"chain"`
	ID           string `bson:"_id"`
	UpdateTime   int64  `bson:"update_time"`
}

type PoolInfo struct {
	Platform           string   `bson:"platform"`              //平台名称
	Chain              string   `bson:"chain"`                 //链名称
	PoolAddress        string   `bson:"pool_address"`          //池子地址    为了查询要存小写
	PoolProperty       string   `bson:"pool_property"`         //池子性质
	IsWatching         string   `bson:"is_watching"`           //是否监控    反射的时候替代不行 所以得字符串
	WatchDescribe      []string `bson:"watch_describe"`        //监控描述
	InputTokenList     []string `bson:"input_token_list"`      //输入token信息列表
	CertificateToken   string   `bson:"certificate_token"`     //凭证token信息
	FundFlowList       []string `bson:"fund_flow_list"`        //资金流向列表
	InputTokenNameList []string `bson:"input_token_name_list"` //输入token名列表主要为了查询  为了查询要存小写
	RewardTokenList    []string `bson:"reward_token_list"`     //奖励token信息列表
	UpdateTime         int64    `bson:"update_time"`           //更新时间
}
