package boolean

import (
	"github.com/antlabs/mock/integer"
)

// 生成一个随机布尔值
func Boolean() bool {
	return integer.IntegerRange(0, 1) == 1
}
