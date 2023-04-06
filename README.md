# mock
生成mock数据，第一个大版本使用反射填充数据。   
[![Go](https://github.com/antlabs/mock/workflows/Go/badge.svg)](https://github.com/antlabs/mock/actions)
[![codecov](https://codecov.io/gh/antlabs/mock/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/mock)
## install
```
go get github.com/antlabs/mock
```
### 用法
```go
type MyType struct {
	Slice []int
	Map   map[string]string
}

type Person struct {
	Name    string
	Age     int
	Address Address
  Email string
}

type Address struct {
	City    string
	Country string
}

type ReferenceType struct {
	MyType      MyType
	Person      Person
	MyTypeP     *MyType
	CreateTime  string
	PointerList []*int
}

var a ReferenceType
mock.MockData(&a)
all, err := json.Marshal(&a)
//输出如下
/*
{
  "Id": "a0050a9f-c476-44bf-afa6-50aa755b412b",
  "MyType": {
    "Slice": [
      1765430006,
      1517363189,
      1513256343,
      1856964067
    ],
    "Map": {
      "0dc39a52cbf7": "1b42ec291a5610",
      "0e96f9fc": "5c8d",
      "0f": "16e68b1c48a0",
      "12f9a60c98": "9ebb",
      "9c19cbc6": "d0a890f22903",
      "c98e58179f33cd29f2": "4ac8cc9d56f84e57cc",
      "d843468e006f75decf": "f84c",
      "f6a851": "4d",
      "fa": "52"
    }
  },
  "Person": {
    "Name": "6ec14adf87637e5e",
    "Age": 1672409902,
    "Address": {
      "City": "9d",
      "Country": "d3"
    }
  },
  "MyTypeP": {
    "Slice": [
      389117540,
      1014359813,
      987503586,
      454448724,
      334452352
    ],
    "Map": {
      "": "8cb115",
      "0345ddcfa73f04229b": "73a82322c31bb569ea",
      "0dff5a973b7bfe62424e": "cb",
      "ab": "b68cdb00dd6a",
      "d57e58a18250b8142939": "cecaab24704788",
      "eaf5cdac7c325943cf": "05532fac46ddaa1198b6"
    }
  },
  "CreateTime": "2035-11-12T04:55:46+08:00",
  "PointerList": [
    948961333,
    1887867813,
    1407499446
  ]
}
*/
```

