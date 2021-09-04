package watch

func BSCWatch() {
	AddPancakePoolInfo("0x62dec3a560d2e8a84d30752ba454f97b26757877", false)
	AddPancakePoolInfo("0x73feaa1ee314f8c655e354234017be2193c9e24e", true)
}

func ETHWatch() {
	//把3pool的每个币种的变化推送
	AddCurvePool3("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7")
}
