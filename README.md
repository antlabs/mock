# mock
[![Go](https://github.com/antlabs/mock/workflows/Go/badge.svg)](https://github.com/antlabs/mock/actions)
[![codecov](https://codecov.io/gh/antlabs/mock/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/mock)

生成mock数据，第一个大版本使用反射填充数据。   
## 支持的类型有
* uint8/uint16/uint32/uint64
* int8/int16/int32/int64
* slice
* map
* time.Time
* 人名
* 国家名
* mac地址
* 省市区
* ...更多

## 一、install
```
go get github.com/antlabs/mock
```
### 二、快速开始
```go
type MyType struct {
	Slice []int
	Map   map[string]string
}

type Person struct {
	Name    string
	Age     int
}

type ReferenceType struct {
	Id          string
	MyType      MyType
	Person      Person
	MyTypeP     *MyType
	CreateTime  string
	PointerList []*int
	Email       string
	URL         string
	UserName    string
	NickName    string
	Country     string
	Ipv4        string
}
var a ReferenceType
mock.MockData(&a)
all, err := json.Marshal(&a)
//输出如下
// {
//   "Person": {
//     "Name": "f9da822c9d",
//     "Age": 48559177
//   },
//   "Id": "6f310b75-914b-4ff7-8b72-eb613100cba3",
//   "MyType": {
//     "Slice": [
//       1376038752,
//       1708184974,
//       1156005062,
//       1349914762,
//       54481841,
//       1714270649,
//       339267240,
//       566602164,
//       367706443,
//       203982404
//     ],
//     "Map": {
//       "55": "f69a1e57a",
//       "602d445c6": "50761",
//       "a94": "959",
//       "c80287f1": "df72f",
//       "d7c4": "60356c4b2"
//     }
//   },
//   "MyTypeP": {
//     "Slice": [
//       1051242817,
//       866106268,
//       1952326595,
//       884632111,
//       587070158,
//       1781740253,
//       844288137,
//       1888080123
//     ],
//     "Map": {
//       "0296d26f": "b",
//       "1d": "133bb8e30c",
//       "3867": "359557",
//       "97379af": "d9",
//       "a6fd02": "a415b9ab08",
//       "a7c9c": "829c45b6",
//       "c855": "59"
//     }
//   },
//   "CreateTime": "2029-05-07T06:28:15+08:00",
//   "PointerList": [
//     303312144,
//     1937065563
//   ],
//   "Email": "cb76@sina.com",
//   "URL": "http://github.com/antlabs/33b6/12/5024/69/b/da0/f8/65de/a6f6/24f/62/e7ec/64/57/54a7/3c/79d/d/16463/b185/101/f8132/d2b/dae/b5b/4/ae7/97ae/8/3/7a6/f92ec/9d/2392/84/9?3=5",
//   "UserName": "士孙玉珍",
//   "NickName": "卜凤玉",
//   "Country": "Cuba",
//   "HeadPic": "www.1.com",
//   "Ipv4": "4.148.97.94",
//   "Province": "湖北省",
//   "City": "恩施土家族苗族自治州",
//   "District": "利川市"
// }
```
## 三、WithXXX各种配置函数
### 3.1 配置指定字段的数据生成范围`WithMinMaxLenByField`
```go
type Test_MinMaxLenByField struct {
	S     string
	Slice []int
}

// 控制slice的生成长度范围
// 控制slice的生成长度范围
func TestMinMaxLenByField(t *testing.T) {
	e := Test_MinMaxLenByField{}
	mock.MockData(&e, mock.WithMinMaxLenByField("S", 10, 20), mock.WithMinMaxLenByField("Slice", 10, 20))
}
```

### 3.2 配置指定字段的数据源`WithContainsFieldSourceString`
指定HeadPic字段的，数据源。
```go
var a ReferenceType
image := []string{"image.xxx.com/1.headpic", "image.xxx.com/2.headpic", "image.xxx.com/3.headpic"}
err := mock.MockData(&a, mock.WithContainsFieldSourceString("HeadPic", image))
```
### 3.3 设置为英文
```go
mock.WithCountryEn()

```

### 3.4 设置数据最大长度`WithMaxLen`
```go
mock.WithMaxLen()
```

### 3.5 设置数据最大长度`WithMinLen`
```go
mock.WithMaxLen()
```

### 3.6 设置数值的最大值`WithMax`
```go
mock.WithMax()
```

### 3.7 设置数值的最大值`WithMin`
```go
mock.WithMin()
```
### 3.8 设置忽略的字段名
字段有时候是由protobuf或者thrift生成，不能直接修改tag，可以使用mock.WithIgnoreFields接口忽略
```go
// 设置忽略的字段名
mock.WithIgnoreFields([]string{"Country", "NickName"})
```
