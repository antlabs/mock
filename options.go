package mock

type Options struct {
	// 生成的mock数据的最大长度
	MaxLen int
	// 生成的mock数据的最小长度
	MinLen int
	// 生成的mock数据的最大值
	Max int
	// 生成的mock数据的最小值
	Min int
}

type Option func(*Options)

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

func WithMax(max int) Option {
	return func(o *Options) {
		o.Max = max
	}
}

func WithMin(min int) Option {
	return func(o *Options) {
		o.Min = min
	}
}
