package stringx

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

const (
	max = 512
	min = 1
)

var bytesPool = sync.Pool{
	New: func() any {
		return make([]byte, max)
	},
}

// 随机生成一个字符串, 范围在[min, max]之间
func StringRange(min, max int) (s string, err error) {
	n := rand.Int31n(int32(max-min+1)) + int32(min)
	p := bytesPool.Get().([]byte)[:n]
	defer bytesPool.Put(p)
	if _, err := rand.Read(p); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", p)[:n], nil

}

// 随机生成一个小写字符串，长度在[min, max]之间
func StringLowerRange(min, max int) (s string, err error) {
	b, err := StringRange(min, max)
	if err != nil {
		return b, err
	}

	return strings.ToLower(b), nil
}

// 随机生成一个字符串，长度在[1, 512]之间
// String returns a random string of length between 1 and 512.
func String() (s string, err error) {
	return StringRange(min, max)
}

// 随机生成一个小写字符串，长度在[1, 512]之间
func Lower() (s string, err error) {

	b, err := String()
	if err != nil {
		return b, err
	}

	return strings.ToLower(b), nil
}

// 随机生成一个大写字符串，长度在[min, max]之间
func StringUpperRange(min, max int) (s string, err error) {
	b, err := StringRange(min, max)
	if err != nil {
		return b, err
	}

	return strings.ToUpper(b), nil
}

// 随机生成一个大写字符串，长度在[1, 512]之间
func Upper() (s string, err error) {

	b, err := String()
	if err != nil {
		return b, err
	}

	return strings.ToUpper(b), nil
}
