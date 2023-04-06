package timex

import (
	"time"

	"github.com/antlabs/mock/integer"
)

type options struct {
	min int64
	max int64
}

type Option func(*options)

func WithMin(min int64) Option {
	return func(o *options) {
		o.min = min
	}
}

func WithMax(max int64) Option {
	return func(o *options) {
		o.max = max
	}
}

func defaultOptions(opts *options) {
	if opts.min == 0 {
		opts.min = time.Now().Unix()
	}

	if opts.max == 0 {
		opts.max = opts.min + 3600*24*365*100
	}
}

// 随便生成一个时间字符串，符合rfc3339格式
// 默认min是当前时间，max是当前时间+100年以内
func TimeRFC3339String(o ...Option) string {
	return TimeRFC3339(o...).Format(time.RFC3339)
}

func TimeRFC3339(o ...Option) time.Time {
	opts := &options{}
	for _, opt := range o {
		opt(opts)
	}

	defaultOptions(opts)

	n := int64(integer.IntegerRangeInt(int(opts.min), int(opts.max)))
	return time.Unix(n, 0)
}

// 随机生成一个时间字符串，格式是yyyy-MM-dd
func YYYYMMDD(o ...Option) string {
	return TimeRFC3339(o...).Format("20060102")
}

// 随机生成一个时间字符串，格式是yyyy
func Year(o ...Option) string {
	opts := &options{}
	for _, opt := range o {
		opt(opts)
	}

	defaultOptions(opts)

	n := int64(integer.IntegerRangeInt(int(opts.min), int(opts.max)))
	return time.Unix(n, 0).Format("2006")
}
