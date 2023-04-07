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
}
var a ReferenceType
mock.MockData(&a)
all, err := json.Marshal(&a)
//输出如下
// {
//   "Id": "fbdce0af-d005-4fc9-8dcd-52e58cad03ea",
//   "MyType": {
//     "Slice": [
//       1616384469,
//       168688132,
//       1888755395,
//       557608762,
//       1971716613,
//       1785522831,
//       1680982268,
//       562918354
//     ],
//     "Map": {
//       "2b5": "82",
//       "2c0": "",
//       "3a4": "8404",
//       "50ed5": "66",
//       "6c05aad": "2d84e",
//       "c7a4a8bc2": "024f70ae7",
//       "d279876e8": "bf68b",
//       "eec4a5845e": ""
//     }
//   },
//   "Person": {
//     "Name": "a672",
//     "Age": 505353573,
//     "Address": {
//       "City": "a48ab3123",
//       "Country": "79153bf"
//     }
//   },
//   "MyTypeP": {
//     "Slice": [
//       2096057502
//     ],
//     "Map": {
//       "25": "8271"
//     }
//   },
//   "CreateTime": "2033-08-15T00:52:03+08:00",
//   "PointerList": [
//     1396718208,
//     1077645781,
//     337823408
//   ],
//   "Email": "f3239e@hotmail.com",
//   "URL": "http://github.com/antlabs/6b165/5774/15/c7bc/29a9/f929/69d/b07/bf469/19/5ea6/e921/10f0/f6c/e57/c5/9a/c4129/e/95e/cb/441/36d/5/f634/7c/3d4de/d/b9/06/bcee7/d3/8fd/b8e/3d24/846/1138/320/37/018/06/433/32773/4313/aec/7a/ae3e0/a26/0d/527af/6c92d/c2a0/ec3ca/8/ded7b/6688c/32/348a9/74b2f/d4320/1/2a/e5/1714/5d99/d/2b/af/ea/7c8/2/1?76ec=8386b",
//   "UserName": "东门玉萍",
//   "NickName": "倚天"
// }
```
## 三、WithXXX各种配置函数
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
