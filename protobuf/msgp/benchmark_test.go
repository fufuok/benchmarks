package msgp

import (
	"testing"

	"github.com/fufuok/bytespool"
)

var message = &Message{
	Query:         "abcdefg",
	PageNumber:    100,
	ResultPerPage: 1,
	Comment:       []byte("xyz"),
	Corpus:        Message_UNIVERSAL,
}

// Benchmark Proto3 Marshal
func BenchmarkProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := message.MarshalMsg(nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkProto3Unmarshal(b *testing.B) {
	data, err := message.MarshalMsg(nil)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var message Message
		_, err := message.UnmarshalMsg(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProto3Marshal_Pool(b *testing.B) {
	var err error
	size := message.Msgsize()
	for i := 0; i < b.N; i++ {
		buf := bytespool.Make(size)
		buf, err = message.MarshalMsg(buf)
		if err != nil {
			b.Fatal(err)
		}
		bytespool.Put(buf)
	}
}
