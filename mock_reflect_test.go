package mock

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type AllTypes struct {
	String   string
	Byte     byte
	Rune     rune
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	Uint     uint
	Uint8    uint8
	Uint16   uint16
	Uint32   uint32
	Uint64   uint64
	Float32  float32
	Float64  float64
	Bool     bool
	Time     time.Time
	Duration time.Duration
	// Complex64  complex64
	// Complex128 complex128
}

// 基础类型
func Test_MockDataBasic(t *testing.T) {
	var a AllTypes
	MockData(&a)
	all, err := json.Marshal(&a)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)

}

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

// 复合类型的测试
func Test_MockData2(t *testing.T) {
	var a ReferenceType
	MockData(&a)
	all, err := json.Marshal(&a)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)
}

type TestEmail struct {
	Email string
}

func Test_MockEmail(t *testing.T) {
	e := TestEmail{}
	MockData(&e)
	all, err := json.Marshal(&e)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)
}

type Test_MinMaxLenByField struct {
	S     string
	Slice []int
}

func TestMinMaxLenByField(t *testing.T) {
	e := Test_MinMaxLenByField{}
	MockData(&e, WithMinMaxLenByField("S", 10, 20), WithMinMaxLenByField("Slice", 10, 20))
	// 检查下生成的长度是否在10-20之间
	assert.True(t, len(e.S) >= 10 && len(e.S) <= 20)
	assert.True(t, len(e.Slice) >= 10 && len(e.S) <= 20)

	all, err := json.Marshal(&e)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)
}
