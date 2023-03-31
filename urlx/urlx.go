package urlx

import (
	"strings"

	"github.com/antlabs/mock/integer"
	"github.com/antlabs/mock/stringx"
)

type options struct {
	protocol string
	domain   string
	path     string
	query    string
}

var PROTOCOLS = []string{
	"http", "https",
}

type Option func(*options)

func WithProtocol(protocol string) Option {
	return func(o *options) {
		o.protocol = protocol
	}
}

func WithDomain(domain string) Option {
	return func(o *options) {
		o.domain = domain
	}
}

func WithPath(path string) Option {
	return func(o *options) {
		o.path = path
	}
}

func WithQuery(query string) Option {
	return func(o *options) {
		o.query = query
	}
}

// 随机生成一个url
func URL(opts ...Option) string {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.protocol == "" {
		o.protocol = Protocol()
	}

	if o.domain == "" {
		o.domain = Domain()
	}

	if o.path == "" {
		o.path = Path()
	}

	if o.query == "" {
		o.query = Query()
	}

	return o.protocol + "://" + o.domain + "/" + o.path + "?" + o.query
}

// 随机生成一个查询字符串
func Query() string {
	k, _ := stringx.StringRange(1, 5)
	v, _ := stringx.StringRange(1, 5)
	return k + "=" + v
}

// 随机生成一个协议
func Protocol() string {
	index := integer.IntegerRange(0, 1)
	return PROTOCOLS[index]
}

// 随机生成一个域名
func Domain() string {
	return "github.com/antlabs"
}

// 随机生成一个路径
func Path() (path string) {
	index := integer.IntegerRange(1, 100)
	var out strings.Builder
	for i := 0; i < index; i++ {
		if i != 0 {
			out.WriteString("/")
		}

		out.WriteString(PathSegment())
	}

	return out.String()
}

// 随机生成一个路径片段
func PathSegment() string {
	s, _ := stringx.StringRange(1, 5)
	return s
}
