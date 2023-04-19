package mock

import (
	"math"
	"reflect"
	"strings"

	"github.com/antlabs/mock/integer"
	"github.com/antlabs/mock/stringx"
	"github.com/antlabs/mock/timex"
)

// 通过反射的方式，对任意类型的数据进行mock
func MockData(x any, opts ...Option) error {
	if x == nil {
		return nil
	}

	opt := &Options{}
	opt.MinLen = 1
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
	// 忽略
	if len(opt.IgnoreFields) > 0 && len(sf.Name) > 0 && opt.IgnoreFields[sf.Name] {
		return nil
	}

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
		l := 0
		if _, ok := opt.MinMaxLenByField[sf.Name]; ok {
			l = integer.IntegerRangeInt(int(opt.MinMaxLenByField[sf.Name].MinLen), int(opt.MinMaxLenByField[sf.Name].MaxLen))
		} else {
			l = integer.IntegerRangeInt(int(opt.MinLen), int(opt.MaxLen))
		}

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

		l := 0
		if _, ok := opt.MinMaxLenByField[sf.Name]; ok {
			l = integer.IntegerRangeInt(int(opt.MinMaxLenByField[sf.Name].MinLen), int(opt.MinMaxLenByField[sf.Name].MaxLen))
		} else {
			l = integer.IntegerRangeInt(int(minLen), int(opt.MaxLen))
		}

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
		var err error
		var ok bool

		var source []string
		// 指定数据来源
		for key, val := range opt.StringSource {
			// TODO 这里是遍历，也需要要优化下
			key = strings.ToLower(key)
			if strings.Contains(strings.ToLower(sf.Name), key) {
				source = val
				goto next
			}
		}

		// 猜测下数据类型
		if err, ok = guessStringType(v, sf, opt); err != nil {
			return err
		} else if !ok {
			return nil
		}

	next:
		if _, ok := opt.MinMaxLenByField[sf.Name]; ok {
			s, err := stringx.StringRange(opt.MinMaxLenByField[sf.Name].MinLen, opt.MinMaxLenByField[sf.Name].MaxLen, stringx.WithSource(source))
			if err != nil {
				return err
			}
			v.SetString(s)
			return nil
		}

		s, err := stringx.StringRange(opt.MinLen, opt.MaxLen, stringx.WithSource(source))
		if err != nil {
			return err
		}

		v.SetString(s)
	}
	return nil
}
