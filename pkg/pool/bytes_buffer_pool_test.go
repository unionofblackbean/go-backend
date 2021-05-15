package pool

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuffer(t *testing.T) {
	//
	// Normal Test
	//
	buf := GetBuffer()
	assert.NotNil(t, buf)
	assert.IsType(t, new(bytes.Buffer), buf)

	//
	// Wrong Type Test
	//
	wrongTypeData := "data with wrong type"
	bytesBufferPool.Put(&wrongTypeData)

	bufFromPool := GetBuffer()
	assert.NotNil(t, bufFromPool)
	assert.IsType(t, new(bytes.Buffer), bufFromPool)
}

func TestPutBuffer(t *testing.T) {
	buf := new(bytes.Buffer)
	PutBuffer(buf)

	bufFromPool := GetBuffer()
	assert.Same(t, buf, bufFromPool)
	assert.IsType(t, new(bytes.Buffer), bufFromPool)
}
