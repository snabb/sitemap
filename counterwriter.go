package sitemap

import (
	"io"
)

// CounterWriter implements [io.Writer]. Number of bytes written is tracked.
type CounterWriter struct {
	writer io.Writer
	count  int64
}

var _ io.Writer = (*CounterWriter)(nil)

// NewCounterWriter wraps [io.Writer] and returns [CounterWriter].
func NewCounterWriter(w io.Writer) (cw *CounterWriter) {
	return &CounterWriter{
		writer: w,
	}
}

// Write calls Write on the wrapped [io.Writer] and adds the number of bytes
// written to the counter.
func (cw *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	cw.count = cw.count + int64(n)
	return n, err
}

// Count returns the number of bytes written to the [CounterWriter].
func (cw *CounterWriter) Count() (n int64) {
	return cw.count
}
