package integer

import (
	"math/rand"
	"sync"
	"time"
)

var intRnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var mux = sync.Mutex{}

// 随机生成一个整数,范围在[min, max]之间
func IntegerRangeInt(min, max int) int {
	// 生成一个随机数
	mux.Lock()
	defer mux.Unlock()

	n := max - min
	return intRnd.Intn(n) + min
}

func IntegerRangeUint(min, max uint) uint {
	// 生成一个随机数
	mux.Lock()
	defer mux.Unlock()

	n := max - min
	return uint(intRnd.Intn(int(n))) + min
}

func Float32Range(min, max float32) float32 {
	mux.Lock()
	defer mux.Unlock()

	return intRnd.Float32()*(max-min) + min
}

func Float64Range(min, max float64) float64 {
	mux.Lock()
	defer mux.Unlock()

	return intRnd.Float64()*(max-min) + min
}

func Float64() float64 {
	mux.Lock()
	defer mux.Unlock()

	return intRnd.Float64()
}

// 随机生成一个整数,范围在[0, 2^61-1]之间
func Integer() int {
	mux.Lock()
	defer mux.Unlock()

	return intRnd.Intn(1<<63 - 1)
}
