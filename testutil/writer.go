package testutil

import (
	"bytes"
	"errors"
)

// TestWriter struct, useful for testing
type TestWriter struct {
	bytes.Buffer
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
	// ErrOnFlush return error on flush, useful for testing
	ErrOnFlush bool
	// ErrOnClose return error on close, useful for testing
	ErrOnClose bool
}

// NewTestWriter instance
func NewTestWriter() *TestWriter {
	return &TestWriter{}
}

// Flush implements
func (w *TestWriter) Flush() error {
	if w.ErrOnFlush {
		return errors.New("flush error")
	}

	w.Reset()
	return nil
}

// Close implements
func (w *TestWriter) Close() error {
	if w.ErrOnClose {
		return errors.New("close error")
	}
	return nil
}

// Write implements
func (w *TestWriter) Write(p []byte) (n int, err error) {
	if w.ErrOnWrite {
		return 0, errors.New("write error")
	}

	return w.Buffer.Write(p)
}
