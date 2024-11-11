package buffer

import "errors"

type ByteBuffer struct {
	bytes []byte
}

func (b *ByteBuffer) Len() int {
	return len(b.bytes)
}

func (b *ByteBuffer) Write(data []byte) (int, error) {
	b.bytes = append(b.bytes, data...)
	return len(b.bytes), nil
}

func (b *ByteBuffer) Read(p []byte) (int, error) {
	if len(b.bytes) == 0 {
		return 0, errors.New("buffer is empty")
	}
	n := copy(p, b.bytes)
	b.bytes = b.bytes[n:]
	return n, nil
}
