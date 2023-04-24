package city

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/antlabs/mock/integer"
)

// 数据引用来源
// https://github.com/modood/Administrative-divisions-of-China/blob/master/dist/pcas-code.json

type cityData []struct {
	Children []struct {
		Children []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"children"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"children"`
	Code string `json:"code"`
	Name string `json:"name"`
}

//go:embed city-code.json
var data []byte

type nameAndCode struct {
	Name     string
	Code     string
	Parent   int // debug使用
	Children []nameAndCode
}

// 所有省
var provinces = []nameAndCode{}

// 所有市
var cities = []nameAndCode{}

// 所有区
var districts = []nameAndCode{}

func init() {
	// 省市区所有的数据
	var allCity cityData
	err := json.Unmarshal(data, &allCity)
	if err != nil {
		panic(err)
	}

	// 对顶级数据进行排序
	sort.Slice(allCity, func(i, j int) bool {
		return allCity[i].Name < allCity[j].Name
	})

	// 所以省份不需要排序
	for _, province := range allCity {
		// 省直接添加
		provinces = append(provinces, nameAndCode{Name: province.Name, Code: province.Code, Parent: 0})
		// 创建一个临时的城市列表
		tmpCities := []nameAndCode{}
		for _, city := range province.Children {
			if city.Name == "市辖区" {
				city.Name = province.Name
			}

			tmpCities = append(tmpCities, nameAndCode{Name: city.Name, Code: city.Code, Parent: len(provinces) - 1})

			// 创建一个临时的区列表
			tmpDistricts := []nameAndCode{}
			for _, district := range city.Children {
				tmpDistricts = append(tmpDistricts, nameAndCode{Name: district.Name, Code: district.Code, Parent: len(cities) - 1})
			}
			tmpCities[len(tmpCities)-1].Children = tmpDistricts

			// 所有的区
			districts = append(districts, tmpDistricts...)
		}

		provinces[len(provinces)-1].Children = tmpCities
		// 拼接到城市列表中
		cities = append(cities, tmpCities...)
	}

	sort.Slice(cities, func(i, j int) bool {
		return cities[i].Name < cities[j].Name
	})
}

func Province() string {
	return provinces[integer.IntegerRangeInt(0, len(provinces)-1)].Name
}

func City(opts ...Option) string {
	var opt Options
	for _, o := range opts {
		o(&opt)
	}

	tempCities := cities
	if opt.ProvinceName != "" {
		data := provinces
		i := sort.Search(len(data), func(i int) bool {
			return data[i].Name >= opt.ProvinceName
		})

		if i < len(data) && data[i].Name == opt.ProvinceName {
			tempCities = data[i].Children
		} else {
			tempCities = cities
		}
	}

	cityItem := tempCities[integer.IntegerRangeInt(0, len(tempCities)-1)]
	return cityItem.Name
}

func District(opts ...Option) string {
	var opt Options
	for _, o := range opts {
		o(&opt)
	}

	tempDistrict := districts
	if opt.CityName != "" {
		data := cities
		i := sort.Search(len(data), func(i int) bool {
			return data[i].Name >= opt.CityName
		})

		if i < len(data) && data[i].Name == opt.CityName {
			tempDistrict = data[i].Children
		} else {
			tempDistrict = districts
		}
	}

	if len(tempDistrict) == 0 {
		panic(fmt.Sprintf("temp district is 0:city name is (%s)", opt.CityName))
	}
	districtsItem := tempDistrict[integer.IntegerRangeInt(0, len(tempDistrict)-1)]

	return districtsItem.Name
}

func ProvinceCode() string {
	return provinces[integer.IntegerRangeInt(0, len(provinces)-1)].Code
}

func CityCode() string {
	return cities[integer.IntegerRangeInt(0, len(provinces)-1)].Code
}

func DistrictCode() string {
	return districts[integer.IntegerRangeInt(0, len(provinces)-1)].Code
}
