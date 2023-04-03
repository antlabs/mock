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

func Test_MockData(t *testing.T) {
	var a AllTypes
	MockData(&a)
	all, err := json.Marshal(&a)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)

}
