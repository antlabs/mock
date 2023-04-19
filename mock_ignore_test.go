package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ignoreCase struct {
	Uint    uint
	Float32 float32
	Float64 float64
	String  string
	Ignored string
}

func TestMockIgnore(t *testing.T) {
	var ignore ignoreCase
	err := MockData(&ignore, WithIgnoreFields("Ignored"))
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), ignore.Uint)
	assert.NotEqual(t, float32(0), ignore.Float32)
	assert.NotEqual(t, float64(0), ignore.Float64)
	assert.NotEqual(t, "", ignore.String)
	assert.Equal(t, "", ignore.Ignored)
}
