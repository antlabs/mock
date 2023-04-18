package mock

import (
	"reflect"
	"strings"

	"github.com/antlabs/mock/country"
	"github.com/antlabs/mock/email"
	"github.com/antlabs/mock/gid"
	"github.com/antlabs/mock/name"
	"github.com/antlabs/mock/timex"
	"github.com/antlabs/mock/urlx"
)

// 猜测string的实际类型
// 第一个返回值是错误, 第二个返回值决定是否继续猜测
func guessStringType(v reflect.Value, sf reflect.StructField, opt *Options) (err error, ok bool) {
	fieldName := sf.Name

	// TODO 优化成map查表函数
	if strings.Contains(strings.ToLower(fieldName), "url") {
		v.SetString(urlx.URL())
		return nil, false
	}

	// 如果字段名是Name，那么就随机生成一个名字
	if strings.Contains(strings.ToLower(fieldName), "username") {
		// TODO 需要修改
		v.SetString(name.Name(name.WithChinese()))
		return nil, false
	}

	// 昵称
	if strings.Contains(strings.ToLower(fieldName), "nickname") {
		v.SetString(name.Name(name.WithChinese()))
		return nil, false
	}

	// 如果字段名是ID，那么就生成一个uuid
	// 忽略大小写搜索id
	if strings.Contains(strings.ToLower(fieldName), "id") {
		v.SetString(gid.GID())
		return nil, false
	}

	// 如果字段名是Time，那么就随机生成一个时间
	if strings.Contains(strings.ToLower(fieldName), "time") {
		v.SetString(timex.TimeRFC3339String(timex.WithMin(opt.Min), timex.WithMax(opt.Max)))
		return nil, false
	}

	// 如果字段名是email，那么就随机生成一个email
	if strings.Contains(strings.ToLower(fieldName), "email") {
		e, err := email.Email()
		if err != nil {
			return err, false
		}
		v.SetString(e)
		return nil, false
	}

	// 如果字段是country, 那么就随机生成一个国家
	if strings.Contains(strings.ToLower(fieldName), "country") {
		v.SetString(country.Country(opt.CountryChina))
		return nil, false
	}

	return nil, true
}
