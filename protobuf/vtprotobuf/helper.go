package pb

import (
	"github.com/fufuok/bytespool"
)

func (m *Message) Marshal(buf []byte) ([]byte, error) {
	size := m.SizeVT()
	if buf == nil {
		buf = bytespool.Get(size)
	}
	n, err := m.MarshalToSizedBufferVT(buf[:size])
	if err != nil {
		return buf[:0], err
	}
	return buf[:n], nil
}

func (m *Message) Unmarshal(buf []byte) error {
	return m.UnmarshalVT(buf)
}
