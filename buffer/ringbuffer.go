package buffer

type RingBuffer struct {
	buffer []byte
	size   int64
}

func NewRingBuffer() *RingBuffer {
	return &RingBuffer{}
}
