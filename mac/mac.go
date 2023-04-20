package mac

import (
	"math/rand"
	"net"
)

// mac 地址
// 目前既支持mac 48格式的生成
func Mac() string {
	mac := make([]byte, 6)
	rand.Read(mac)
	// net.HardwareAddr比fmt.Sprintf更快
	// return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
	return net.HardwareAddr(mac).String()
}
