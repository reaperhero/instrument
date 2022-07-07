package commslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntsHas(t *testing.T) {
	ints := []int{2, 4, 5}
	assert.True(t, IntsHas(ints, 2))
	assert.False(t, IntsHas(ints, 3))
}

func TestInt64sHas(t *testing.T) {
	ints := []int64{2, 4, 5}
	assert.True(t, Int64sHas(ints, 2))
	assert.True(t, Int64sHas(ints, 5))
	assert.False(t, Int64sHas(ints, 3))
}
