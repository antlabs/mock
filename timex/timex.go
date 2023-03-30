package timex

import (
	"time"

	"github.com/antlabs/mock/integer"
)

// 随便生成一个时间字符串，符合rfc3339格式
func Time() string {
	n := integer.IntegerRange(0, 1<<63-1)
	return time.Unix(int64(n), 0).Format(time.RFC3339)
}
