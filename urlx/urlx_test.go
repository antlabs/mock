package urlx

import (
	"fmt"
	"testing"
)

// 自测URL接口，生成一个随机的url
func TestURL(t *testing.T) {
	s := URL()
	fmt.Println(s)
}
