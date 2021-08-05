package data

import (
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/rand"
	"sync"
	"time"
)

type WrapReplyList struct {
	Mu        *sync.RWMutex
	ReplyList []*analysis.HA_Single_Reply
}

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
