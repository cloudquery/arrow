package memory

import (
	"github.com/apache/arrow/go/v13/arrow/internal/debug"
)

func AssertBuffer(reader string, buffer *Buffer) {
	if buffer == nil {
		return
	}
	debug.Assert(buffer.refCount == 1, reader+":buffer.refCount")
}
