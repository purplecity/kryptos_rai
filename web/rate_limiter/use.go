package rate_limiter

import (
	"net/http"
	"strings"
)

var u_rules Rules
var u_ratelimit RateLimiter

func init() {
	u_rules = NewRules()
	u_ratelimit, _ = NewRateLimiter("memory")
}

func RateLimitAddRules(pattern string, duration, limit int32) {
	u_rules.AddRule(pattern, &Rule{Duration: duration, Limit: limit})
}

func RateLimitInitRules() {
	u_ratelimit.InitRules(u_rules)
}

func RateLimitCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIPPort := r.RemoteAddr
		s := strings.Split(clientIPPort, ":")
		var ok = true
		if len(s) != 2 {
			ok = false
		}
		if !ok {
			w.WriteHeader(http.StatusForbidden) //403
			return
		}

		pattern := r.RequestURI
		ok = u_ratelimit.TokenAccess(s[0], pattern)

		if !ok {
			w.WriteHeader(http.StatusTooManyRequests) //429
			return
		}
		next.ServeHTTP(w, r)
	})
}
