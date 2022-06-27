package commslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringItem(t *testing.T) {
	source := []string{"1", "2", "3", "4","4"}
	dst := RemoveSliceStringItem(source, "4")
	assert.Equal(t, dst, []string{"1", "2", "3"})
}
