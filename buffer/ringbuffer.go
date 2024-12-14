package buffer

type RingBuffer struct {
	buffer        []byte
	size          int
	startPosition int
	endPosition   int
}

func NewRingBuffer() *RingBuffer {
	return &RingBuffer{}
}
