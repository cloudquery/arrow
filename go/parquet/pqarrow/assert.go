package pqarrow

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/parquet/internal/debug"
)

func assertBuildArray(reader string, arr *arrow.Chunked, err error) {
	if err != nil {
		return
	}
	debug.Assert(arr != nil, reader+":assertBuildArray: arr != nil")
	if arr == nil {
		return
	}
	for _, chunk := range arr.Chunks() {
		array.AssertArray(reader, chunk)
	}
}
