package ipv4

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func validateIPv4(ip string) bool {
	// 解析IPv4地址
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// 检查地址格式是否正确
	if parsedIP.To4() == nil {
		return false
	}

	return true
}

func TestIPv4(t *testing.T) {
	i4 := IPv4()
	assert.True(t, validateIPv4(i4))
}
