package ipv4

import (
	"fmt"

	"github.com/antlabs/mock/integer"
)

func IPv4() string {
	// 生成4个随机数
	var octets [4]uint16
	for i := 0; i < 4; i++ {
		octets[i] = uint16(integer.IntegerRangeInt(0, 255))
	}

	// 返回IPv4地址
	return fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3])
}
