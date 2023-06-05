package array

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/internal/debug"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func AssertArray(reader string, arr arrow.Array) {
	debug.Assert(arr != nil, reader+":AssertArray: arr != nil")
	if arr == nil {
		return
	}
	assertData(reader, arr.Data().(*Data))
}

func assertData(reader string, data *Data) {
	debug.Assert(data != nil, reader+":data != nil")
	if data == nil {
		return
	}
	debug.Assert(data.refCount == 1, reader+":data.refCount == 1")
	for _, buff := range data.buffers {
		memory.AssertBuffer(reader, buff)
	}
	if data.dictionary != nil {
		assertData(reader, data.dictionary)
	}
	for _, child := range data.childData {
		assertData(reader, child.(*Data))
	}
}
