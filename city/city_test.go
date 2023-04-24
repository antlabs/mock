package city

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_City_guangdong(t *testing.T) {
	s := City(WithProvinceName("广东省"))
	v := slices.Contains([]string{"广州市", "深圳市", "珠海市", "汕头市", "佛山市", "韶关市", "湛江市", "肇庆市", "江门市", "茂名市", "惠州市", "梅州市", "汕尾市", "河源市", "阳江市", "清远市", "东莞市", "中山市", "潮州市", "揭阳市", "云浮市"}, s)
	assert.True(t, v)
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

func Test_City_notfound(t *testing.T) {
	s := City(WithProvinceName("上海市1"))
	assert.NotEmpty(t, s)
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

// 测试区
// 重庆市
func Test_City_chongqing_district(t *testing.T) {
	s := District(WithCityName("重庆市"))
	v := slices.Contains([]string{"万州区", "涪陵区", "渝中区", "大渡口区", "江北区", "沙坪坝区", "九龙坡区", "南岸区", "北碚区", "綦江区", "大足区", "渝北区", "巴南区", "黔江区", "长寿区", "江津区", "合川区", "永川区", "南川区", "璧山区", "铜梁区", "潼南区", "荣昌区", "开州区", "梁平区", "武隆区", "城口县", "丰都县", "垫江县", "忠县", "云阳县", "奉节县", "巫山县", "巫溪县", "石柱土家族自治县", "秀山土家族苗族自治县", "酉阳土家族苗族自治县", "彭水苗族土家族自治县"}, s)
	assert.True(t, v, fmt.Sprintf("get %s", s))
}
