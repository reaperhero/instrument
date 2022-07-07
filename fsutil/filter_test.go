package fsutil_test

import (
	"github.com/reaperhero/instrument/fsutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFunc(t *testing.T) {
	var fn fsutil.FilterFunc = func(filePath, filename string) bool { return false }

	assert.False(t, fn("path/some.txt", "some.txt"))
}
