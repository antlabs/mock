package mac

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMac(t *testing.T) {
	mac := Mac()
	_, err := net.ParseMAC(mac)
	assert.NoError(t, err)
}
