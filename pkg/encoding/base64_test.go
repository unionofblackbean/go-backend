package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64RawStdEncodeToString(t *testing.T) {
	assert.Equal(t, "SGVsbG8gV29ybGQh", Base64RawStdEncodeToString([]byte("Hello World!")))
}

func BenchmarkBase64RawStdEncodeToString(b *testing.B) {
	data := []byte("Hello World!")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64RawStdEncodeToString(data)
	}
}

func TestBase64RawStdDecodeString(t *testing.T) {
	result, err := Base64RawStdDecodeString("SGVsbG8gV29ybGQh")
	assert.Nil(t, err)

	assert.Equal(t, []byte("Hello World!"), result)
}

func BenchmarkBase64RawStdDecodeString(b *testing.B) {
	data := "Hello World!"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64RawStdDecodeString(data)
	}
}
