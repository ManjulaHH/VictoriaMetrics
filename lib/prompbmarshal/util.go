package prompbmarshal

import (
	"fmt"
)

// MarshalWriteRequest marshals wr to dst and returns the result.
func MarshalWriteRequest(dst []byte, wr *WriteRequest) []byte {
	size := wr.Size()
	dstLen := len(dst)
	if n := size - (cap(dst) - dstLen); n > 0 {
		dst = append(dst[:cap(dst)], make([]byte, n)...)
	}
	dst = dst[:dstLen+size]
	n, err := wr.MarshalToSizedBuffer(dst[dstLen:])
	if err != nil {
		panic(fmt.Errorf("BUG: unexpected error when marshaling WriteRequest: %s", err))
	}
	return dst[:dstLen+n]
}

// ResetWriteRequest resets wr.
func ResetWriteRequest(wr *WriteRequest) {
	for i := range wr.Timeseries {
		ts := wr.Timeseries[i]
		ts.Labels = nil
		ts.Samples = nil
	}
	wr.Timeseries = wr.Timeseries[:0]
}
