package string

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const max = 512

var bytesPool = sync.Pool{
	New: func() any {
		return make([]byte, max)
	},
}

// go随机生成一段string字符串
// String returns a random string of length between 1 and 512.
// TODO 优化下面的代码
func String() (s string, err error) {
	n := rand.Int31n(max)
	p := bytesPool.Get().([]byte)[:n]
	defer bytesPool.Put(p)
	if _, err := rand.Read(p); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", p), nil
}

func Lower() (s string, err error) {

	b, err := String()
	if err != nil {
		return b, err
	}

	return strings.ToLower(b), nil
}

func Upper() (s string, err error) {

	b, err := String()
	if err != nil {
		return b, err
	}

	return strings.ToUpper(b), nil
}
