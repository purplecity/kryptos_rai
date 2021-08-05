package watch

import (
	"math/rand"
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
