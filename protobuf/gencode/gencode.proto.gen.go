package gencode

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Message struct {
	Query         string
	PageNumber    int32
	ResultPerPage int32
	Comment       []byte
	Corpus        int32
}

func (d *Message) Size() (s uint64) {

	{
		l := uint64(len(d.Query))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Comment))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 12
	return
}
func (d *Message) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Query))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Query)
		i += l
	}
	{

		buf[i+0+0] = byte(d.PageNumber >> 0)

		buf[i+1+0] = byte(d.PageNumber >> 8)

		buf[i+2+0] = byte(d.PageNumber >> 16)

		buf[i+3+0] = byte(d.PageNumber >> 24)

	}
	{

		buf[i+0+4] = byte(d.ResultPerPage >> 0)

		buf[i+1+4] = byte(d.ResultPerPage >> 8)

		buf[i+2+4] = byte(d.ResultPerPage >> 16)

		buf[i+3+4] = byte(d.ResultPerPage >> 24)

	}
	{
		l := uint64(len(d.Comment))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Comment)
		i += l
	}
	{

		buf[i+0+8] = byte(d.Corpus >> 0)

		buf[i+1+8] = byte(d.Corpus >> 8)

		buf[i+2+8] = byte(d.Corpus >> 16)

		buf[i+3+8] = byte(d.Corpus >> 24)

	}
	return buf[:i+12], nil
}

func (d *Message) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Query = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		d.PageNumber = 0 | (int32(buf[i+0+0]) << 0) | (int32(buf[i+1+0]) << 8) | (int32(buf[i+2+0]) << 16) | (int32(buf[i+3+0]) << 24)

	}
	{

		d.ResultPerPage = 0 | (int32(buf[i+0+4]) << 0) | (int32(buf[i+1+4]) << 8) | (int32(buf[i+2+4]) << 16) | (int32(buf[i+3+4]) << 24)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Comment)) >= l {
			d.Comment = d.Comment[:l]
		} else {
			d.Comment = make([]byte, l)
		}
		copy(d.Comment, buf[i+8:])
		i += l
	}
	{

		d.Corpus = 0 | (int32(buf[i+0+8]) << 0) | (int32(buf[i+1+8]) << 8) | (int32(buf[i+2+8]) << 16) | (int32(buf[i+3+8]) << 24)

	}
	return i + 12, nil
}
