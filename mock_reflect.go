package mock

import (
	"math"
	"reflect"
	"strings"

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
	return mockData(reflect.ValueOf(x), opt)
}

func defaultOptions(opt *Options) {
	if opt.Max == 0 {
		opt.Max = math.MaxInt32
	}

	if opt.MaxLen == 0 {
		opt.MaxLen = 10
	}
}

func mockData(v reflect.Value, opt *Options) error {
	switch v.Kind() {
	// 指针类型，需要先获取指针指向的值
	case reflect.Ptr:
		return mockData(v.Elem(), opt)
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

			if err := mockData(v.Field(i), opt); err != nil {
				return err
			}

		}

		// slice 或者 array 类型，需要遍历所有元素
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 && reflect.Array == v.Kind() {
			return nil
		}

		for i := 0; i < v.Len(); i++ {
			if err := mockData(v.Index(i), opt); err != nil {
				return err
			}
		}

		// map类型，需要遍历map的所有key
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := mockData(v.MapIndex(key), opt); err != nil {
				return err
			}
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
		// 获取字段名
		name := v.Type().Name()
		// 如果字段名是ID，那么就生成一个uuid
		// 忽略大小写搜索id
		if strings.Contains(strings.ToLower(name), "id") {
			v.SetString(gid.GID())
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
