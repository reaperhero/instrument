package commslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringItem(t *testing.T) {
	source := []string{"1", "2", "3", "4","4"}
	dst := StringsRemove(source, "4")
	assert.Equal(t, dst, []string{"1", "2", "3"})
}



func TestStringsHas(t *testing.T) {
	ss := []string{"a", "b"}
	assert.True(t, StringsHas(ss, "a"))
	assert.True(t, StringsHas(ss, "b"))
	assert.False(t, StringsHas(ss, "c"))
}
