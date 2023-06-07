package array

import (
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/internal/debug"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func AssertArray(reader string, arr arrow.Array) {
	debug.Assert(arr != nil, reader+":AssertArray: arr != nil")
	if arr == nil {
		return
	}
	AssertData(reader, arr.Data().(*Data))
}

func AssertData(reader string, data *Data) {
	debug.Assert(data != nil, reader+":data != nil")
	if data == nil {
		return
	}
	debug.Assert(data.refCount == 1, reader+":data.refCount="+strconv.Itoa(int(data.refCount)))
	for i, buff := range data.buffers {
		memory.AssertBuffer(reader+"["+strconv.Itoa(i)+"]", buff)
	}
	if data.dictionary != nil {
		AssertData(reader, data.dictionary)
	}
	for _, child := range data.childData {
		AssertData(reader, child.(*Data))
	}
}
