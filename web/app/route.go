package app

import (
	"kryptos_rai/web/rate_limiter"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

var routes []Route

func init() {
	register("POST", "/web/HandlePoolInfo", HandlePoolInfo, 1, 5)   //后台池子操作
	register("POST", "/web/AddressAnalysis", AddressAnalysis, 1, 5) //地址余额查询
	register("POST", "/web/HandleLogoInfo", HandleLogoInfo, 1, 5)   // 后台logo操作
	register("POST", "/web/PoolComboBox", PoolComboBox, 1, 5)       //后台池子下拉框选项
	register("POST", "/web/LogoComboBox", LogoComboBox, 1, 5)       //后台logo下拉框选项
	register("POST", "/web/HandleSubcribe", HandleSubcribe, 1, 5)   //订阅
	rate_limiter.RateLimitInitRules()
}

//在duration 秒内 允许的limit请求次数
func register(method, pattern string, handler http.HandlerFunc, duration, limit int32) {
	routes = append(routes, Route{method, pattern, handler})
	rate_limiter.RateLimitAddRules(pattern, duration, limit)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)
	}
	router.Use(rate_limiter.RateLimitCheck)
	return router
}
