package mock

// minLen 和 maxLen 用于指定生成的mock数据的最大长度和最小长度
type MinMax struct {
	// 生成的mock数据的最大长度
	MaxLen int
	// 生成的mock数据的最小长度
	MinLen int
}

type Options struct {
	// 指定字段名，指定生成的数据来源
	StringSource map[string][]string
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
	// 默认是mock
	TagName string
	// 设置忽略的字段名
	IgnoreFields map[string]bool
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

// 设置最大长度
func WithMaxLen(maxLen int) Option {
	return func(o *Options) {
		o.MaxLen = maxLen
	}
}

// 设置最小长度
func WithMinLen(minLen int) Option {
	return func(o *Options) {
		o.MinLen = minLen
	}
}

// 设置最大值
func WithMax(max int64) Option {
	return func(o *Options) {
		o.Max = max
	}
}

// 设置最小值
func WithMin(min int64) Option {
	return func(o *Options) {
		o.Min = min
	}
}

// 设置country为英文
func WithCountryEn() Option {
	return func(o *Options) {
		o.CountryChina = true
	}
}

// 包含field字符串，就会使用source中的数据, 字段比较是string类型
func WithContainsFieldSourceString(field string, source []string) Option {
	return func(o *Options) {
		if o.StringSource == nil {
			o.StringSource = make(map[string][]string)
		}

		o.StringSource[field] = source
	}
}

// 设置tag名
// TODO实现
func WithTag(tag string) Option {
	return func(o *Options) {
		o.TagName = tag
	}
}

// 设置忽略的字段名
func WithIgnoreFields(fields ...string) Option {
	return func(o *Options) {
		if o.IgnoreFields == nil {
			o.IgnoreFields = make(map[string]bool)
		}
		for _, s := range fields {
			o.IgnoreFields[s] = true
		}
	}
}
