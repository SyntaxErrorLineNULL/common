package buffer

type ByteBuffer struct {
	Bytes []byte
}

func (b *ByteBuffer) Len() int {
	return len(b.Bytes)
}
