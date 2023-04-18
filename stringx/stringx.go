package stringx

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/antlabs/mock/integer"
)

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
func StringRange(min, max int, opts ...Option) (s string, err error) {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	n := integer.IntegerRangeInt(min, max)
	p := bytesPool.Get().([]byte)[:n]
	defer bytesPool.Put(p)

	// 指定数据来源
	if len(o.Source) > 0 {

		pos := integer.IntegerRangeInt(0, len(o.Source)-1)

		s = o.Source[pos]
		if len(s) > max {
			s = s[:max]
		}

	} else {
		// 随机生成
		if _, err := rand.Read(p); err != nil {
			return "", err
		}
		s = fmt.Sprintf("%x", p)[:n]
	}

	if o.ToUpper {
		return strings.ToUpper(s), nil
	}

	if o.ToUpper {
		return strings.ToLower(s), nil
	}
	return s, nil
}
