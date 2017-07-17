package msg

type OffHookMessage struct {
	LineInstance  uint32
	CallReference uint32
}

func (m *OffHookMessage) Id() uint32 {
	return OffHookMessageId
}

func (m *OffHookMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallReference)
}

func (m *OffHookMessage) Deserialize(d *Deserializer) {
	m.LineInstance = d.ReadUint32()
	m.CallReference = d.ReadUint32()
}

func init() {
	registerFactory(
		OffHookMessageId,
		func() Msg { return &OffHookMessage{} },
	)
}

type SetSpeakerModeMessage struct {
	SpeakerMode uint32
}

func (m *SetSpeakerModeMessage) Id() uint32 {
	return OffHookMessageId
}

func (m *SetSpeakerModeMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.SpeakerMode)
}

func (m *SetSpeakerModeMessage) Deserialize(d *Deserializer) {
	m.SpeakerMode = d.ReadUint32()
}

func init() {
	registerFactory(
		SetSpeakerModeMessageId,
		func() Msg { return &SetSpeakerModeMessage{} },
	)
}

type DefineTimeDateMessage struct {
	TimeDateInfo [32]byte
	SystemTime   uint32
}

func (m *DefineTimeDateMessage) Id() uint32 {
	return DefineTimeDateMessageId
}

func (m *DefineTimeDateMessage) Serialize(s *Serializer) {
	s.WriteBytes(m.TimeDateInfo[:])
	s.WriteUint32(m.SystemTime)
}

func (m *DefineTimeDateMessage) Deserialize(d *Deserializer) {
	d.ReadBytes(m.TimeDateInfo[:])
	m.SystemTime = d.ReadUint32()
}

func init() {
	registerFactory(
		DefineTimeDateMessageId,
		func() Msg { return &DefineTimeDateMessage{} },
	)
}
