package mock

import (
	"math"
	"reflect"
	"strings"

	"github.com/antlabs/mock/email"
	"github.com/antlabs/mock/gid"
	"github.com/antlabs/mock/integer"
	"github.com/antlabs/mock/stringx"
	"github.com/antlabs/mock/timex"
)

// 通过反射的方式，对任意类型的数据进行mock
func MockData(x any, opts ...Option) error {
	if x == nil {
		return nil
	}

	var opt = &Options{}
	for _, o := range opts {
		o(opt)
	}

	defaultOptions(opt)
	return mockData(reflect.ValueOf(x), reflect.StructField{}, opt)
}

func defaultOptions(opt *Options) {
	if opt.Max == 0 {
		opt.Max = math.MaxInt32
	}

	if opt.MaxLen == 0 {
		opt.MaxLen = 10
	}
}

func mockData(v reflect.Value, sf reflect.StructField, opt *Options) error {
	switch v.Kind() {
	// 指针类型，需要先获取指针指向的值
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		return mockData(v.Elem(), sf, opt)
		// 结构体类型，需要遍历结构体的所有字段
	case reflect.Struct:
		// 如果是time.Time类型，直接返回当前时间
		if v.Type().String() == "time.Time" {
			tv := timex.TimeRFC3339(timex.WithMin(opt.Min), timex.WithMax(opt.Max))
			v.Set(reflect.ValueOf(tv))
			return nil
		}
		typ := v.Type()
		for i := 0; i < v.NumField(); i++ {
			sf := typ.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous {
				continue
			}

			if err := mockData(v.Field(i), sf, opt); err != nil {
				return err
			}

		}

		// slice 或者 array 类型，需要遍历所有元素
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 && reflect.Array == v.Kind() {
			return nil
		}

		// 随机生成一个长度
		l := integer.IntegerRangeInt(int(opt.MinLen), int(opt.MaxLen))
		// 如果是slice类型，那么就需要扩容
		if reflect.Slice == v.Kind() {
			v.Set(reflect.MakeSlice(v.Type(), l, l))
		}

		for i := 0; i < v.Len(); i++ {
			if err := mockData(v.Index(i), sf, opt); err != nil {
				return err
			}
		}

		// map类型，需要遍历map的所有key
	case reflect.Map:
		if v.Len() > 0 {
			return nil
		}

		// 随机生成map的长度
		minLen := opt.MinLen
		if minLen == 0 {
			minLen = 1
		}

		l := integer.IntegerRangeInt(int(minLen), int(opt.MaxLen))
		// 创建一个map
		v.Set(reflect.MakeMapWithSize(v.Type(), l))
		// 遍历map的所有key

		for i := 0; i < l; i++ {
			// 创建一个key
			key := reflect.New(v.Type().Key()).Elem()
			// 创建一个value
			value := reflect.New(v.Type().Elem()).Elem()
			// 递归mock key
			if err := mockData(key, sf, opt); err != nil {
				return err
			}
			// 递归mock value
			if err := mockData(value, sf, opt); err != nil {
				return err
			}
			// 设置map的key和value
			v.SetMapIndex(key, value)
		}

		// 接口类型，需要先获取接口的值
	case reflect.Interface:
		// float32, float64 类型
	case reflect.Float32, reflect.Float64:
		f := integer.Float64Range(float64(opt.Min), float64(opt.Max))
		v.SetFloat(f)
		// int... 类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := integer.IntegerRangeInt(int(opt.Min), int(opt.Max))
		v.SetInt(int64(i))
		// uint... 类型
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := integer.IntegerRangeUint(uint(opt.Min), uint(opt.Max))

		v.SetUint(uint64(u))
		// string 类型
	case reflect.String:
		// 获取字段变量名
		name := sf.Name

		// 如果字段名是ID，那么就生成一个uuid
		// 忽略大小写搜索id
		if strings.Contains(strings.ToLower(name), "id") {
			v.SetString(gid.GID())
			return nil
		}

		// 如果字段名是Time，那么就随机生成一个时间
		if strings.Contains(strings.ToLower(name), "time") {
			v.SetString(timex.TimeRFC3339String(timex.WithMin(opt.Min), timex.WithMax(opt.Max)))
			return nil
		}

		// 如果字段名是email，那么就随机生成一个email
		if strings.Contains(strings.ToLower(name), "email") {
			e, err := email.Email()
			if err != nil {
				return err
			}
			v.SetString(e)
			return nil
		}
		s, err := stringx.StringRange(opt.MinLen, opt.MaxLen)
		if err != nil {
			return err
		}

		v.SetString(s)
	}
	return nil
}
