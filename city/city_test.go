package city

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_City(t *testing.T) {
}

func Test_City_jiangsu(t *testing.T) {
	s := City(WithProvinceName("江苏省"))
	v := slices.Contains([]string{"南京市", "苏州市", "无锡市", "常州市", "镇江市", "南通市", "泰州市", "扬州市", "盐城市", "连云港市", "徐州市", "淮安市", "宿迁市"}, s)
	assert.True(t, v)
}

func Test_City_shanghai(t *testing.T) {
	s := City(WithProvinceName("上海市"))
	v := slices.Contains([]string{"上海市"}, s)
	assert.True(t, v)
}

// 测试区
func Test_City_shanghai_district(t *testing.T) {
	s := District(WithCityName("上海市"))
	v := slices.Contains([]string{"黄浦区", "徐汇区", "长宁区", "静安区", "普陀区", "虹口区", "杨浦区", "闵行区", "宝山区", "嘉定区", "浦东新区", "金山区", "松江区", "青浦区", "奉贤区", "崇明区"}, s)
	assert.True(t, v, fmt.Sprintf("get %s", s))
}

func Test_City_yancheng(t *testing.T) {
	s := District(WithCityName("盐城市"))
	v := slices.Contains([]string{"亭湖区", "盐都区", "大丰区", "响水县", "滨海县", "阜宁县", "射阳县", "建湖县", "东台市", "盐城经济技术开发区"}, s)
	assert.True(t, v, fmt.Sprintf("get %s", s))
}
