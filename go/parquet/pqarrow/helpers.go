package pqarrow

import (
	"github.com/apache/arrow/go/v13/arrow"
)

func releaseArrData(data ...arrow.ArrayData) {
	for _, d := range data {
		if d != nil {
			d.Release()
		}
	}
}
