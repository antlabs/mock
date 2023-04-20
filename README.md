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
	Address Address
}

type Address struct {
	City    string
	Country string
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
//   "Id": "45b00376-ebf0-4196-9c95-6a8aa25c88c0",
//   "MyType": {
//     "Slice": [
//       1400477113,
//       1591736785,
//       720060954,
//       259088419,
//       1728262781,
//       1452739972,
//       1171863610,
//       2007929494,
//       346720360
//     ],
//     "Map": {
//       "012f4ecd": "5ad0378fc1",
//       "01415828fb": "bcb4e437",
//       "1cf8ef3": "a0189d",
//       "2": "301a2e456d",
//       "4012ad12": "6f4cee5",
//       "4d5ca43": "088c3",
//       "d": "58",
//       "e87e3f": "ecc601a"
//     }
//   },
//   "Person": {
//     "Name": "c9cc8ae",
//     "Age": 383225028,
//     "Address": {
//       "City": "d2700",
//       "Country": "St. Kitts & Nevis"
//     }
//   },
//   "MyTypeP": {
//     "Slice": [
//       1692231780,
//       942566115,
//       1429602187,
//       370936121,
//       450946004
//     ],
//     "Map": {
//       "395e33e6": "5e3",
//       "46470d59c": "608ca0e8",
//       "804e4": "2",
//       "a": "40f751",
//       "b0dd": "933",
//       "eb5a28060a": "4ac6728",
//       "f2": "84d6bf9ece"
//     }
//   },
//   "CreateTime": "2024-09-20T18:45:02+08:00",
//   "PointerList": [
//     1717968446,
//     1265293376
//   ],
//   "Email": "74fea6@hotmail.com",
//   "URL": "https://github.com/antlabs/30/0/ce1/ee16/0/c07f/3a67/25/dc3/6/cae/4/cc/2b9ce/10bd1/e9059/33/31a18/7/6/bc/b?0495=5a",
//   "UserName": "檀娜",
//   "NickName": "全志强",
//   "Country": "Gibraltar",
//   "HeadPic": "www.3.com",
//   "Ipv4": "29.40.97.24"
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
