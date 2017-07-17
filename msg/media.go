package msg

type MediaPathEventMessage struct {
	MediaPathID    uint32
	MediaPathEvent uint32
}

func (m *MediaPathEventMessage) Id() uint32 {
	return MediaPathEventMessageId
}

func (m *MediaPathEventMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.MediaPathID)
	s.WriteUint32(m.MediaPathEvent)
}

func (m *MediaPathEventMessage) Deserialize(d *Deserializer) {
	m.MediaPathID = d.ReadUint32()
	m.MediaPathEvent = d.ReadUint32()
}

func init() {
	registerFactory(
		MediaPathEventMessageId,
		func() Msg { return &MediaPathEventMessage{} },
	)
}

type OpenReceiveChannelMessage struct {
	ConferenceID              uint32
	PassThruPartyID           uint32
	MillisecondPacketSize     uint32
	CompressionType           uint32
	QualifierIn               [8]byte
	CallReference             uint32
	MRxMediaEncryptionKeyInfo [48]byte
	StreamPassThroughID       uint32
	AssociatedStreamID        uint32
	RFC2833PayloadType        uint32
	DtmfType                  uint32
	MixingMode                uint32
	PartyDirection            uint32
	Ipv4or6                   uint32
	SourceIpAddr              [1]byte
	SourcePortNumber          [4]byte
	RequestedIpAddrType       uint32
	AudioLevelAdjustment      uint32
	LatentCapsInfo            [36]byte
}

func (m *OpenReceiveChannelMessage) Id() uint32 {
	return OpenReceiveChannelMessageId
}

func (m *OpenReceiveChannelMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.ConferenceID)
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.MillisecondPacketSize)
	s.WriteUint32(m.CompressionType)
	s.WriteBytes(m.QualifierIn[:])
	s.WriteUint32(m.CallReference)
	s.WriteBytes(m.MRxMediaEncryptionKeyInfo[:])
	s.WriteUint32(m.StreamPassThroughID)
	s.WriteUint32(m.AssociatedStreamID)
	s.WriteUint32(m.RFC2833PayloadType)
	s.WriteUint32(m.DtmfType)
	s.WriteUint32(m.MixingMode)
	s.WriteUint32(m.PartyDirection)
	s.WriteUint32(m.Ipv4or6)
	s.WriteBytes(m.SourceIpAddr[:])
	s.WriteBytes(m.SourcePortNumber[:])
	s.WriteUint32(m.RequestedIpAddrType)
	s.WriteUint32(m.AudioLevelAdjustment)
	s.WriteBytes(m.LatentCapsInfo[:])
}

func (m *OpenReceiveChannelMessage) Deserialize(d *Deserializer) {
	m.ConferenceID = d.ReadUint32()
	m.PassThruPartyID = d.ReadUint32()
	m.MillisecondPacketSize = d.ReadUint32()
	m.CompressionType = d.ReadUint32()
	d.ReadBytes(m.QualifierIn[:])
	m.CallReference = d.ReadUint32()
	d.ReadBytes(m.MRxMediaEncryptionKeyInfo[:])
	m.StreamPassThroughID = d.ReadUint32()
	m.AssociatedStreamID = d.ReadUint32()
	m.RFC2833PayloadType = d.ReadUint32()
	m.DtmfType = d.ReadUint32()
	m.MixingMode = d.ReadUint32()
	m.PartyDirection = d.ReadUint32()
	m.Ipv4or6 = d.ReadUint32()
	d.ReadBytes(m.SourceIpAddr[:])
	d.ReadBytes(m.SourcePortNumber[:])
	m.RequestedIpAddrType = d.ReadUint32()
	m.AudioLevelAdjustment = d.ReadUint32()
	d.ReadBytes(m.LatentCapsInfo[:])
}

func init() {
	registerFactory(
		OpenReceiveChannelMessageId,
		func() Msg { return &OpenReceiveChannelMessage{} },
	)
}

type OpenReceiveChannelAckMessage struct {
	OpenReceiveChannelStatus uint32
	Ipv4or6                  uint32
	IpAddr                   [16]byte
	PortNumber               [4]byte
	PassThruPartyID          uint32
	CallReference            uint32
}

func (m *OpenReceiveChannelAckMessage) Id() uint32 {
	return OpenReceiveChannelAckMessageId
}

func (m *OpenReceiveChannelAckMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.OpenReceiveChannelStatus)
	s.WriteUint32(m.Ipv4or6)
	s.WriteBytes(m.IpAddr[:])
	s.WriteBytes(m.PortNumber[:])
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.CallReference)
}

func (m *OpenReceiveChannelAckMessage) Deserialize(d *Deserializer) {
	m.OpenReceiveChannelStatus = d.ReadUint32()
	m.Ipv4or6 = d.ReadUint32()
	d.ReadBytes(m.IpAddr[:])
	d.ReadBytes(m.PortNumber[:])
	m.PassThruPartyID = d.ReadUint32()
	m.CallReference = d.ReadUint32()
}

func init() {
	registerFactory(
		OpenReceiveChannelAckMessageId,
		func() Msg { return &OpenReceiveChannelAckMessage{} },
	)
}

type CloseReceiveChannelMessage struct {
	ConferenceID     uint32
	PassThruPartyID  uint32
	CallReference    uint32
	PortHandlingFlag uint32
}

func (m *CloseReceiveChannelMessage) Id() uint32 {
	return CloseReceiveChannelMessageId
}

func (m *CloseReceiveChannelMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.ConferenceID)
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.CallReference)
	s.WriteUint32(m.PortHandlingFlag)
}

func (m *CloseReceiveChannelMessage) Deserialize(d *Deserializer) {
	m.ConferenceID = d.ReadUint32()
	m.PassThruPartyID = d.ReadUint32()
	m.CallReference = d.ReadUint32()
	m.PortHandlingFlag = d.ReadUint32()
}

func init() {
	registerFactory(
		CloseReceiveChannelMessageId,
		func() Msg { return &CloseReceiveChannelMessage{} },
	)
}

type StartMediaTransmitionMessage struct {
	ConferenceID              uint32
	PassThruPartyID           uint32
	Ipv4or6                   uint32
	RemoteIpAddr              [16]byte
	RemotePortNumber          uint32
	MillisecondPacketSize     uint32
	CompressionType           uint32
	QualifierOut              [16]byte
	CallReference             uint32
	MTxMediaEncryptionKeyInfo [48]byte
	StreamPassThroughID       uint32
	AssociatedStreamID        uint32
	RFC2833PayloadType        uint32
	DtmfType                  uint32
	MixingMode                uint32
	PartyDirection            uint32
	LatentCapsInfo            [36]byte
}

func (m *StartMediaTransmitionMessage) Id() uint32 {
	return StartMediaTransmitionMessageId
}

func (m *StartMediaTransmitionMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.ConferenceID)
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.Ipv4or6)
	s.WriteBytes(m.RemoteIpAddr[:])
	s.WriteUint32(m.RemotePortNumber)
	s.WriteUint32(m.MillisecondPacketSize)
	s.WriteUint32(m.CompressionType)
	s.WriteBytes(m.QualifierOut[:])
	s.WriteUint32(m.CallReference)
	s.WriteBytes(m.MTxMediaEncryptionKeyInfo[:])
	s.WriteUint32(m.StreamPassThroughID)
	s.WriteUint32(m.AssociatedStreamID)
	s.WriteUint32(m.RFC2833PayloadType)
	s.WriteUint32(m.DtmfType)
	s.WriteUint32(m.MixingMode)
	s.WriteUint32(m.PartyDirection)
	s.WriteBytes(m.LatentCapsInfo[:])
}

func (m *StartMediaTransmitionMessage) Deserialize(d *Deserializer) {
	m.ConferenceID = d.ReadUint32()
	m.PassThruPartyID = d.ReadUint32()
	m.Ipv4or6 = d.ReadUint32()
	d.ReadBytes(m.RemoteIpAddr[:])
	m.RemotePortNumber = d.ReadUint32()
	m.MillisecondPacketSize = d.ReadUint32()
	m.CompressionType = d.ReadUint32()
	d.ReadBytes(m.QualifierOut[:])
	m.CallReference = d.ReadUint32()
	d.ReadBytes(m.MTxMediaEncryptionKeyInfo[:])
	m.StreamPassThroughID = d.ReadUint32()
	m.AssociatedStreamID = d.ReadUint32()
	m.RFC2833PayloadType = d.ReadUint32()
	m.DtmfType = d.ReadUint32()
	m.MixingMode = d.ReadUint32()
	m.PartyDirection = d.ReadUint32()
	d.ReadBytes(m.LatentCapsInfo[:])
}

func init() {
	registerFactory(
		StartMediaTransmitionMessageId,
		func() Msg { return &StartMediaTransmitionMessage{} },
	)
}

type StartMediaTransmitionAckMessage struct {
	ConferenceID                uint32
	PassThruPartyID             uint32
	CallReference               uint32
	Ipv4or6                     uint32
	RemoteIpAddr                [16]byte
	RemotePortNumber            [4]byte
	StartMediaTransmitionStatus uint32
}

func (m *StartMediaTransmitionAckMessage) Id() uint32 {
	return StartMediaTransmitionAckMessageId
}

func (m *StartMediaTransmitionAckMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.ConferenceID)
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.CallReference)
	s.WriteUint32(m.Ipv4or6)
	s.WriteBytes(m.RemoteIpAddr[:])
	s.WriteBytes(m.RemotePortNumber[:])
	s.WriteUint32(m.StartMediaTransmitionStatus)
}

func (m *StartMediaTransmitionAckMessage) Deserialize(d *Deserializer) {
	m.ConferenceID = d.ReadUint32()
	m.PassThruPartyID = d.ReadUint32()
	m.CallReference = d.ReadUint32()
	m.Ipv4or6 = d.ReadUint32()
	d.ReadBytes(m.RemoteIpAddr[:])
	d.ReadBytes(m.RemotePortNumber[:])
	m.StartMediaTransmitionStatus = d.ReadUint32()
}

func init() {
	registerFactory(
		StartMediaTransmitionAckMessageId,
		func() Msg { return &StartMediaTransmitionAckMessage{} },
	)
}

type StopMediaTransmitionMessage struct {
	ConferenceID     uint32
	PassThruPartyID  uint32
	CallReference    uint32
	PortHandlingFlag uint32
}

func (m *StopMediaTransmitionMessage) Id() uint32 {
	return StopMediaTransmitionMessageId
}

func (m *StopMediaTransmitionMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.ConferenceID)
	s.WriteUint32(m.PassThruPartyID)
	s.WriteUint32(m.CallReference)
	s.WriteUint32(m.PortHandlingFlag)
}

func (m *StopMediaTransmitionMessage) Deserialize(d *Deserializer) {
	m.ConferenceID = d.ReadUint32()
	m.PassThruPartyID = d.ReadUint32()
	m.CallReference = d.ReadUint32()
	m.PortHandlingFlag = d.ReadUint32()
}

func init() {
	registerFactory(
		StopMediaTransmitionMessageId,
		func() Msg { return &StopMediaTransmitionMessage{} },
	)
}
