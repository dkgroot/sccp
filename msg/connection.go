package msg

type ConnectionStatisticsReqMessage struct {
	DirectoryNum        [25]byte
	CallReference       uint32
	StatsProcessingMode uint32
}

func (m *ConnectionStatisticsReqMessage) Id() uint32 {
	return ConnectionStatisticsReqMessageId
}

func (m *ConnectionStatisticsReqMessage) Serialize(s *Serializer) {
	s.WriteBytes(m.DirectoryNum[:])
	s.WriteUint32(m.CallReference)
	s.WriteUint32(m.StatsProcessingMode)
}

func (m *ConnectionStatisticsReqMessage) Deserialize(d *Deserializer) {
	d.ReadBytes(m.DirectoryNum[:])
	m.CallReference = d.ReadUint32()
	m.StatsProcessingMode = d.ReadUint32()
}

func init() {
	registerFactory(
		ConnectionStatisticsReqMessageId,
		func() Msg { return &ConnectionStatisticsReqMessage{} },
	)
}
