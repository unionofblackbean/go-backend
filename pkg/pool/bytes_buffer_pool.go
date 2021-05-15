package pool

import (
	"bytes"
	"sync"
)

var bytesBufferPool sync.Pool

func GetBuffer() *bytes.Buffer {
	if rawBytesBuffer := bytesBufferPool.Get(); rawBytesBuffer != nil {
		buf, ok := rawBytesBuffer.(*bytes.Buffer)
		if ok {
			buf.Reset()
			return buf
		}
	}

	return new(bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}
