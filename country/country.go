package country

import (
	"embed"
	_ "embed"
	"encoding/csv"
	"io"

	"github.com/antlabs/mock/integer"
)

// TODO: 数据来源，以及清洗脚本, ci自动化
type country struct {
	EnName string
	ZhName string
}

//go:embed country.dat
var data embed.FS

var countrySlice = []country{}

func init() {
	fd, err := data.Open("country.dat")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// 创建一个CSV reader
	reader := csv.NewReader(fd)
	reader.Comma = ','

	for {
		// 读取一行记录
		record, err := reader.Read()
		// 如果到达文件末尾，则退出循环
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(record) != 3 {
			continue
		}

		countrySlice = append(countrySlice, country{EnName: record[1], ZhName: record[2]})
	}
}

func Country(china bool) string {
	pos := integer.IntegerRangeInt(0, len(countrySlice)-1)
	if china {
		return countrySlice[pos].ZhName
	}

	return countrySlice[pos].EnName
}
