package msg

import (
	"errors"
	"io"
)

const (
	headerSize  = 12
	maxBodySize = 2000
)

type MsgWriter struct {
	w   io.Writer
	buf [headerSize + maxBodySize]byte
}

func NewMsgWriter(w io.Writer) *MsgWriter {
	return &MsgWriter{w: w}
}

func (mw *MsgWriter) WriteMsg(m Msg) error {
	// serialize body
	serializer := &Serializer{Buf: mw.buf[headerSize:]}
	m.Serialize(serializer)
	bodySize := serializer.Off

	// serialize header
	serializer = &Serializer{Buf: mw.buf[:headerSize]}
	serializer.WriteUint32(uint32(bodySize) + 4)
	serializer.WriteUint32(0)
	serializer.WriteUint32(m.Id())

	// write msg
	_, err := mw.w.Write(mw.buf[:headerSize+bodySize])

	return err
}

type MsgReader struct {
	r   io.Reader
	buf [maxBodySize]byte
}

func NewMsgReader(r io.Reader) *MsgReader {
	return &MsgReader{r: r}
}

func (mr *MsgReader) ReadMsg() (Msg, error) {
	var m Msg

	// read header
	n, err := mr.r.Read(mr.buf[:headerSize])
	if err != nil {
		return m, err
	} else if n != headerSize {
		return m, errors.New("could read header completly")
	}

	// deserialize header
	deserializer := &Deserializer{Buf: mr.buf[:headerSize]}
	length := deserializer.ReadUint32()
	deserializer.ReadUint32()
	msgId := deserializer.ReadUint32()

	bodySize := int(length) - 4
	if bodySize < 0 || bodySize > maxBodySize {
		return m, errors.New("invalid message lenght")
	}

	// read body
	_, err = io.ReadFull(mr.r, mr.buf[:bodySize])
	if err != nil {
		return m, err
	}

	// deserialize body
	deserializer = &Deserializer{Buf: mr.buf[:bodySize]}
	m = newFromId(msgId)
	m.Deserialize(deserializer)

	return m, nil
}
