package sysutil_test

import (
	"os"
	"testing"

	"github.com/reaperhero/instrument/sysutil"
	"github.com/stretchr/testify/assert"
)

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
