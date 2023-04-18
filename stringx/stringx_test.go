package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	s, err := StringRange(10, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(s))
}
