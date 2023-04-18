# mock
[![Go](https://github.com/antlabs/mock/workflows/Go/badge.svg)](https://github.com/antlabs/mock/actions)
[![codecov](https://codecov.io/gh/antlabs/mock/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/mock)

生成mock数据，第一个大版本使用反射填充数据。   
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
}
var a ReferenceType
mock.MockData(&a)
all, err := json.Marshal(&a)
//输出如下

// {
//   "Id": "2db4f77a-ada2-47cd-9fa5-2d3d7d759d18",
//   "MyType": {
//     "Slice": [
//       1681352613,
//       20691731,
//       681305396,
//       1674447608,
//       266532997,
//       1292383940
//     ],
//     "Map": {
//       "2": "afcb",
//       "4": "f2",
//       "569": "fd2",
//       "a": "b59b48",
//       "a8792": "cee",
//       "bdaf35": "0d19dfe4",
//       "be2a50b4": "f19b56d8"
//     }
//   },
//   "Person": {
//     "Name": "ac0a61a4",
//     "Age": 337830828,
//     "Address": {
//       "City": "",
//       "Country": "U.S. Outlying Islands"
//     }
//   },
//   "MyTypeP": {
//     "Slice": [
//       1807852194,
//       2141777709,
//       92514607,
//       1286243933
//     ],
//     "Map": {
//       "34": "82ab6bf60",
//       "5fa6d": "b",
//       "636484814d": "",
//       "77": "da",
//       "82": "407e",
//       "85784638": "",
//       "a11b8a68": "00c9403dc",
//       "e28c": "b37424e"
//     }
//   },
//   "CreateTime": "2034-11-29T08:17:35+08:00",
//   "PointerList": [
//     895889511,
//     69415086,
//     952093756,
//     381246645,
//     775471733,
//     121055351,
//     2032148785,
//     697090480
//   ],
//   "Email": "c15a950@qq.com",
//   "URL": "http://github.com/antlabs/a0c/6d/4/d021/297/1bd/d1/b2dc/f25?3e=3",
//   "UserName": "寇舒舒",
//   "NickName": "谏静雯",
//   "Country": "Norway"
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
