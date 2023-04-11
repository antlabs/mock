package mock

// minLen 和 maxLen 用于指定生成的mock数据的最大长度和最小长度
type MinMax struct {
	// 生成的mock数据的最大长度
	MaxLen int
	// 生成的mock数据的最小长度
	MinLen int
}

type Options struct {
	// 生成的mock数据的最大长度
	MaxLen int
	// 生成的mock数据的最小长度
	MinLen int
	// 生成的mock数据的最大值
	Max int64
	// 生成的mock数据的最小值
	Min int64
	// 指定字段名，生成最大长度
	MinMaxLenByField map[string]MinMax
	// country 用中文还是英文
	CountryChina bool
}

type Option func(*Options)

// 指定字段名，指定生成长度的范围, slice的长度和string的长度
func WithMinMaxLenByField(field string, minLen int, maxLen int) Option {
	return func(o *Options) {
		if o.MinMaxLenByField == nil {
			o.MinMaxLenByField = make(map[string]MinMax)
		}

		o.MinMaxLenByField[field] = MinMax{MaxLen: maxLen, MinLen: minLen}
	}
}

func WithMaxLen(maxLen int) Option {
	return func(o *Options) {
		o.MaxLen = maxLen
	}
}

func WithMinLen(minLen int) Option {
	return func(o *Options) {
		o.MinLen = minLen
	}
}

func WithMax(max int64) Option {
	return func(o *Options) {
		o.Max = max
	}
}

func WithMin(min int64) Option {
	return func(o *Options) {
		o.Min = min
	}
}

// 设置country为中文
func WithCountryEn() Option {
	return func(o *Options) {
		o.CountryChina = true
	}
}
