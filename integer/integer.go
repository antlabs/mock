package integer

import (
	"math/rand"
	"sync"
	"time"
)

var intRnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var mux = sync.Mutex{}

// 随机生成一个整数,范围在[min, max]之间
func IntegerRange(min, max int) int {
	// 生成一个随机数
	mux.Lock()
	defer mux.Unlock()

	return intRnd.Intn(max-min+1) + min
}

// 随机生成一个整数,范围在[0, 2^61-1]之间
func Integer() int {
	mux.Lock()
	defer mux.Unlock()

	intRnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	return intRnd.Intn(1<<63 - 1)
}
