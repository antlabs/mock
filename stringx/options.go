package stringx

type Options struct {
	Source []string
	// 大写
	ToUpper bool
	// 小写
	ToLower bool
}

type Option func(*Options)

// 设置数据来源
func WithSource(s []string) Option {
	return func(o *Options) {
		o.Source = s
	}
}

func WithUpper() Option {
	return func(o *Options) {
		o.ToUpper = true
	}
}

func WithLower() Option {
	return func(o *Options) {
		o.ToLower = true
	}
}
