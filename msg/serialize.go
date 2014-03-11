package msg

import (
	"encoding/binary"
)

type Serializer struct {
	Buf []byte
	Off int
}

func (s *Serializer) WriteUint8(v uint8) {
	s.Buf[s.Off] = v
	s.Off += 1
}

func (s *Serializer) WriteUint32(v uint32) {
	binary.LittleEndian.PutUint32(s.Buf[s.Off:], v)
	s.Off += 4
}

func (s *Serializer) WriteBytes(v []byte) {
	n := copy(s.Buf[s.Off:], v)
	if n != len(v) {
		panic("couldn't copy all bytes from v")
	}
	s.Off += n
}

type Deserializer struct {
	Buf []byte
	Off int
}

func (d *Deserializer) ReadUint8() uint8 {
	v := d.Buf[d.Off]
	d.Off += 1
	return v
}

func (d *Deserializer) ReadUint32() uint32 {
	v := binary.LittleEndian.Uint32(d.Buf[d.Off:])
	d.Off += 4
	return v
}

func (d *Deserializer) ReadBytes(v []byte) {
	n := copy(v, d.Buf[d.Off:])
	if n != len(v) {
		panic("couldn't copy all bytes to v")
	}
	d.Off += n
}
