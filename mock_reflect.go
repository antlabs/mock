package mock

import (
	"reflect"

	"github.com/antlabs/mock/integer"
	"github.com/antlabs/mock/stringx"
)

// 通过反射的方式，对任意类型的数据进行mock
func MockData(x any, opts ...Option) error {
	if x == nil {
		return nil
	}

	var opt = &Options{}
	return mockData(reflect.ValueOf(x), opt)
}

func mockData(v reflect.Value, opt *Options) error {
	switch v.Kind() {
	case reflect.Ptr:
		return mockData(v.Elem(), opt)
	case reflect.Struct:
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

	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {

		}
	case reflect.Map:
	case reflect.Interface:
	case reflect.Float32, reflect.Float64:
		f := integer.Float64Range(float64(opt.Min), float64(opt.Max))
		v.SetFloat(f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := integer.IntegerRangeInt(opt.Min, opt.Max)
		v.SetInt(int64(i))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := integer.IntegerRangeUint(uint(opt.Min), uint(opt.Max))

		v.SetUint(uint64(u))
	case reflect.String:
		s, err := stringx.StringRange(opt.Min, opt.Max)
		if err != nil {
			return err
		}

		v.SetString(s)
	}
	return nil
}
