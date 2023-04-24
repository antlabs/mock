package city

type Option func(*Options)

type Options struct {
	// 省份名称
	ProvinceName string

	// 城市名称
	CityName string
	// 区名称
	// District string
}

func WithProvinceName(name string) Option {
	return func(o *Options) {
		o.ProvinceName = name
	}
}

func WithCityName(name string) Option {
	return func(o *Options) {
		o.CityName = name
	}
}
