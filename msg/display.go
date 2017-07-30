package msg

type DisplayPromptStatusV2Message struct {
	TimeOutValue uint32
	LineInstance uint32
	CallInstance uint32
	PromptStatus uint32
}

func (m *DisplayPromptStatusV2Message) Id() uint32 {
	return DisplayPromptStatusV2MessageId
}

func (m *DisplayPromptStatusV2Message) Serialize(s *Serializer) {
	s.WriteUint32(m.TimeOutValue)
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallInstance)
	s.WriteUint32(m.PromptStatus)
}

func (m *DisplayPromptStatusV2Message) Deserialize(d *Deserializer) {
	m.TimeOutValue = d.ReadUint32()
	m.LineInstance = d.ReadUint32()
	m.CallInstance = d.ReadUint32()
	m.PromptStatus = d.ReadUint32()
}

func init() {
	registerFactory(
		DisplayPromptStatusV2MessageId,
		func() Msg { return &DisplayPromptStatusV2Message{} },
	)
}

type DisplayPriNotifyV2Message struct {
	TimeOutValue uint32
	Priority     uint32
	Notify       [12]byte
}

func (m *DisplayPriNotifyV2Message) Id() uint32 {
	return DisplayPriNotifyV2MessageId
}

func (m *DisplayPriNotifyV2Message) Serialize(s *Serializer) {
	s.WriteUint32(m.TimeOutValue)
	s.WriteUint32(m.Priority)
	s.WriteBytes(m.Notify[:])
}

func (m *DisplayPriNotifyV2Message) Deserialize(d *Deserializer) {
	m.TimeOutValue = d.ReadUint32()
	m.Priority = d.ReadUint32()
	d.ReadBytes(m.Notify[:])
}

func init() {
	registerFactory(
		DisplayPriNotifyV2MessageId,
		func() Msg { return &DisplayPriNotifyV2Message{} },
	)
}

type ClearPromptStatusMessage struct {
	LineInstance  uint32
	CallReference uint32
}

func (m *ClearPromptStatusMessage) Id() uint32 {
	return DisplayPriNotifyV2MessageId
}

func (m *ClearPromptStatusMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallReference)
}

func (m *ClearPromptStatusMessage) Deserialize(d *Deserializer) {
	m.LineInstance = d.ReadUint32()
	m.CallReference = d.ReadUint32()
}

func init() {
	registerFactory(
		ClearPromptStatusMessageId,
		func() Msg { return &ClearPromptStatusMessage{} },
	)
}
