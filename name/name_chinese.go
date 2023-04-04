package name

import (
	_ "embed"
	"strings"

	"github.com/antlabs/mock/integer"
)

type Options struct {
	// 是否中文
	IsChinese bool
}

var xings []string
var nans []string
var nvs []string

func init() {
	xings = strings.Split(xing, "\n")
	nans = strings.Split(nan, "\n")
	nvs = strings.Split(nv, "\n")
}

//go:embed xing.dat
var xing string

//go:embed nan.dat
var nan string

//go:embed nv.dat
var nv string

func WithChinese() Options {
	return Options{
		IsChinese: true,
	}
}

func CName() string {
	// 先生成姓
	// 然后生成名
	mings := nans
	if integer.IntegerRangeInt(0, 1) == 0 {
		mings = nvs
	}

	xingPos := integer.IntegerRangeInt(0, len(xings))
	mingPos := integer.IntegerRangeInt(0, len(mings))
	return xings[xingPos] + mings[mingPos]
}
