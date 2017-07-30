package msg

type SetLampMessage struct {
	Stimulus          uint32
	StimmulusInstance uint32
	LampMode          uint32
}

func (m *SetLampMessage) Id() uint32 {
	return SetLampMessageId
}

func (m *SetLampMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.Stimulus)
	s.WriteUint32(m.StimmulusInstance)
	s.WriteUint32(m.LampMode)
}

func (m *SetLampMessage) Deserialize(d *Deserializer) {
	m.Stimulus = d.ReadUint32()
	m.StimmulusInstance = d.ReadUint32()
	m.LampMode = d.ReadUint32()
}

func init() {
	registerFactory(
		SetLampMessageId,
		func() Msg { return &SetLampMessage{} },
	)
}
