package commslice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFormatter(t *testing.T) {
	arr := [2]string{"a", "b"}
	str := FormatIndent(arr, "  ")
	assert.Contains(t, str, "\n  ")
	fmt.Println(str)

	str = FormatIndent(arr, "")
	assert.NotContains(t, str, "\n  ")
	assert.Equal(t, "[a, b]", str)
	fmt.Println(str)

	assert.Equal(t, "", FormatIndent("invalid", ""))
	assert.Equal(t, "[]", FormatIndent([]string{}, ""))
}
