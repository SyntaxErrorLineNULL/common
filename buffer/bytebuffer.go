package buffer

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
