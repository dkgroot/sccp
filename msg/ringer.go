package msg

type SetRingerMessage struct {
	RingMode      uint32
	RingDuration  uint32
	LineInstance  uint32
	CallReference uint32
}

func (m *SetRingerMessage) Id() uint32 {
	return SetRingerMessageId
}

func (m *SetRingerMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.RingMode)
	s.WriteUint32(m.RingDuration)
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallReference)
}

func (m *SetRingerMessage) Deserialize(d *Deserializer) {
	m.RingMode = d.ReadUint32()
	m.RingDuration = d.ReadUint32()
	m.LineInstance = d.ReadUint32()
	m.CallReference = d.ReadUint32()
}

func init() {
	registerFactory(
		SetRingerMessageId,
		func() Msg { return &SetRingerMessage{} },
	)
}
