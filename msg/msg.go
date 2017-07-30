package msg

const (
	// Messages sent from client to server
	KeepAliveMsgId                    = 0x00
	RegisterMsgId                     = 0x01
	KeypadButtonMsgId                 = 0x03
	OffHookMessageId                  = 0x06
	OpenReceiveChannelAckMessageId    = 0x22
	MediaPathEventMessageId           = 0x49
	SetSpeakerModeMessageId           = 0x88
	StartMediaTransmitionAckMessageId = 0x154

	// Messages sent from server to client
	StartMediaTransmitionMessageId   = 0x8A
	StopMediaTransmitionMessageId    = 0x8B
	RegisterAckMsgId                 = 0x81
	DefineTimeDateMessageId          = 0x94
	CapabilitiesReqMsgId             = 0x9B
	RegisterRejMsgId                 = 0x9D
	KeepAliveAckMsgId                = 0x0100
	OpenReceiveChannelMessageId      = 0x0105
	CloseReceiveChannelMessageId     = 0x0106
	ConnectionStatisticsReqMessageId = 0x0107
	SelectSoftKeysMessageId          = 0x0110
	ClearPromptStatusMessageId       = 0x0113
	ClearNotifyMsgId                 = 0x0115
	ActivateCallPlaneMessageId       = 0x0116
	CallStateMessageId               = 0x0111
	DisplayPriNotifyV2MessageId      = 0x0144
	DisplayPromptStatusV2MessageId   = 0x0145
	CallInfoV2MessageId              = 0x014A
	StopToneMessageId                = 0x83
	SetLampMessageId                 = 0x86
	SetRingerMessageId               = 0x85

	// Device types (used in RegisterMsg)
	DeviceType7960   = 7
	DeviceType7940   = 8
	DeviceType7941   = 115
	DeviceType7911   = 307
	DeviceType7941GE = 309
	DeviceType7921   = 365
	DeviceType7906   = 369
	DeviceType7962   = 404
	DeviceType7937   = 431
	DeviceType7942   = 434
	DeviceType7905   = 20000
	DeviceType7970   = 30006
	DeviceType7912   = 30007
	DeviceType7961   = 30018
)

type Msg interface {
	Id() uint32

	// Serialize the message, excluding the ID, into s.
	//
	// If the serializer buffer is too small, the function must panic (meh)
	Serialize(s *Serializer)

	// Unserialize the message, excluding the ID, into s.
	//
	// If the deserializer buffer is too small, the function must panic (meh)
	Deserialize(d *Deserializer)
}

// Special case, represent a message that is not mapped to a real
// type otherwise
type GenericMsg struct {
	MsgId uint32
	Data  []byte
}

func (m *GenericMsg) Id() uint32 {
	return m.MsgId
}

func (m *GenericMsg) Serialize(s *Serializer) {
	s.WriteBytes(m.Data)
}

func (m *GenericMsg) Deserialize(d *Deserializer) {
	m.Data = make([]byte, len(d.Buf))
	d.ReadBytes(m.Data)
}

type KeepAliveMsg struct{}

func (m *KeepAliveMsg) Id() uint32 {
	return KeepAliveMsgId
}

func (m *KeepAliveMsg) Serialize(s *Serializer) {
}

func (m *KeepAliveMsg) Deserialize(d *Deserializer) {
}

func init() {
	registerFactory(
		KeepAliveMsgId,
		func() Msg { return &KeepAliveMsg{} },
	)
}

type RegisterMsg struct {
	Name          [16]byte
	UserId        uint32
	LineInstance  uint32
	Ip            uint32
	Type          uint32
	MaxStreams    uint32
	ActiveStreams uint32
	ProtoVersion  uint8
}

func (m *RegisterMsg) Id() uint32 {
	return RegisterMsgId
}

func (m *RegisterMsg) Serialize(s *Serializer) {
	s.WriteBytes(m.Name[:])
	s.WriteUint32(m.UserId)
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.Ip)
	s.WriteUint32(m.Type)
	s.WriteUint32(m.MaxStreams)
	s.WriteUint32(m.ActiveStreams)
	s.WriteUint8(m.ProtoVersion)
}

func (m *RegisterMsg) Deserialize(d *Deserializer) {
	d.ReadBytes(m.Name[:])
	m.UserId = d.ReadUint32()
	m.LineInstance = d.ReadUint32()
	m.Ip = d.ReadUint32()
	m.Type = d.ReadUint32()
	m.MaxStreams = d.ReadUint32()
	m.ActiveStreams = d.ReadUint32()
	m.ProtoVersion = d.ReadUint8()
}

func init() {
	registerFactory(
		RegisterMsgId,
		func() Msg { return &RegisterMsg{} },
	)
}

type StopToneMessage struct {
	LineInstance  uint32
	CallReference uint32
	Tone          uint32
}

func (m *StopToneMessage) Id() uint32 {
	return StopToneMessageId
}

func (m *StopToneMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallReference)
	s.WriteUint32(m.Tone)
}

func (m *StopToneMessage) Deserialize(d *Deserializer) {
	m.LineInstance = d.ReadUint32()
	m.CallReference = d.ReadUint32()
	m.Tone = d.ReadUint32()
}

func init() {
	registerFactory(
		StopToneMessageId,
		func() Msg { return &StopToneMessage{} },
	)
}

type SelectSoftKeysMessage struct {
	LineInstance    uint32
	CallInstance    uint32
	SoftKeySetIndex uint32
	ValidKeyMask    [4]byte
}

func (m *SelectSoftKeysMessage) Id() uint32 {
	return SelectSoftKeysMessageId
}

func (m *SelectSoftKeysMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallInstance)
	s.WriteUint32(m.SoftKeySetIndex)
	s.WriteBytes(m.ValidKeyMask[:])
}

func (m *SelectSoftKeysMessage) Deserialize(d *Deserializer) {
	m.LineInstance = d.ReadUint32()
	m.CallInstance = d.ReadUint32()
	m.SoftKeySetIndex = d.ReadUint32()
	d.ReadBytes(m.ValidKeyMask[:])
}

func init() {
	registerFactory(
		SelectSoftKeysMessageId,
		func() Msg { return &SelectSoftKeysMessage{} },
	)
}

type CallStateMessage struct {
	CallState    uint32
	LineInstance uint32
	CallInstance uint32
}

func (m *CallStateMessage) Id() uint32 {
	return CallStateMessageId
}

func (m *CallStateMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.CallState)
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallInstance)
}

func (m *CallStateMessage) Deserialize(d *Deserializer) {
	m.CallState = d.ReadUint32()
	m.LineInstance = d.ReadUint32()
	m.CallInstance = d.ReadUint32()
}

func init() {
	registerFactory(
		CallStateMessageId,
		func() Msg { return &CallStateMessage{} },
	)
}

type KeypadButtonMsg struct {
	Button       uint32
	LineInstance uint32
	CallInstance uint32
}

func (m *KeypadButtonMsg) Id() uint32 {
	return KeypadButtonMsgId
}

func (m *KeypadButtonMsg) Serialize(s *Serializer) {
	s.WriteUint32(m.Button)
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallInstance)
}

func (m *KeypadButtonMsg) Deserialize(d *Deserializer) {
	m.Button = d.ReadUint32()
	m.LineInstance = d.ReadUint32()
	m.CallInstance = d.ReadUint32()
}

func init() {
	registerFactory(
		KeypadButtonMsgId,
		func() Msg { return &KeypadButtonMsg{} },
	)
}

type CallInfoV2Message struct {
	LineInstance               uint32
	CallReference              uint32
	CallType                   uint32
	OriginalCdpnRedirectReason uint32
	LastRedirectingReason      uint32
	CallInstance               uint32
	CallSecurityStatus         uint32
	PartyPIRestrictionBits     uint32
	CallingParty               [10]byte
	AlternateCallingParty      [10]byte
	CalledParty                [8]byte
	OriginalCalledParty        [8]byte
	LastRedirectingParty       [8]byte
	CallingPartyName           [14]byte
	// OriginalCalledPartyName
	// LastRedirectingPartyName
}

func (m *CallInfoV2Message) Id() uint32 {
	return CallInfoV2MessageId
}

func (m *CallInfoV2Message) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstance)
	s.WriteUint32(m.CallInstance)
	s.WriteUint32(m.CallType)
	s.WriteUint32(m.OriginalCdpnRedirectReason)
	s.WriteUint32(m.LastRedirectingReason)
	s.WriteUint32(m.CallInstance)
	s.WriteUint32(m.CallSecurityStatus)
	s.WriteUint32(m.PartyPIRestrictionBits)
	s.WriteBytes(m.CallingParty[:])
	s.WriteBytes(m.AlternateCallingParty[:])
	s.WriteBytes(m.CalledParty[:])
	s.WriteBytes(m.OriginalCalledParty[:])
	s.WriteBytes(m.LastRedirectingParty[:])
	s.WriteBytes(m.CallingPartyName[:])
}

func (m *CallInfoV2Message) Deserialize(d *Deserializer) {
	m.LineInstance = d.ReadUint32()
	m.CallReference = d.ReadUint32()
	m.CallType = d.ReadUint32()
	m.OriginalCdpnRedirectReason = d.ReadUint32()
	m.LastRedirectingReason = d.ReadUint32()
	m.CallInstance = d.ReadUint32()
	m.CallSecurityStatus = d.ReadUint32()
	m.PartyPIRestrictionBits = d.ReadUint32()
	d.ReadBytes(m.CallingParty[:])
	d.ReadBytes(m.AlternateCallingParty[:])
	d.ReadBytes(m.CalledParty[:])
	d.ReadBytes(m.OriginalCalledParty[:])
	d.ReadBytes(m.LastRedirectingParty[:])
	d.ReadBytes(m.CallingPartyName[:])
}

func init() {
	registerFactory(
		CallInfoV2MessageId,
		func() Msg { return &CallInfoV2Message{} },
	)
}

type RegisterAckMsg struct {
	KeepAlive          uint32
	DateTemplate       [6]byte
	Res                [2]byte
	SecondaryKeepAlive uint32
	ProtoVersion       uint8
	Unknown1           uint8
	Unknown2           uint8
	Unknown3           uint8
}

func (m *RegisterAckMsg) Id() uint32 {
	return RegisterAckMsgId
}

func (m *RegisterAckMsg) Serialize(s *Serializer) {
	s.WriteUint32(m.KeepAlive)
	s.WriteBytes(m.DateTemplate[:])
	s.WriteBytes(m.Res[:])
	s.WriteUint32(m.SecondaryKeepAlive)
	s.WriteUint8(m.ProtoVersion)
	s.WriteUint8(m.Unknown1)
	s.WriteUint8(m.Unknown2)
	s.WriteUint8(m.Unknown3)
}

func (m *RegisterAckMsg) Deserialize(d *Deserializer) {
	m.KeepAlive = d.ReadUint32()
	d.ReadBytes(m.DateTemplate[:])
	d.ReadBytes(m.Res[:])
	m.SecondaryKeepAlive = d.ReadUint32()
	m.ProtoVersion = d.ReadUint8()
	m.Unknown1 = d.ReadUint8()
	m.Unknown2 = d.ReadUint8()
	m.Unknown3 = d.ReadUint8()
}

func init() {
	registerFactory(
		RegisterAckMsgId,
		func() Msg { return &RegisterAckMsg{} },
	)
}

type KeepAliveAckMsg struct{}

func (m *KeepAliveAckMsg) Id() uint32 {
	return KeepAliveAckMsgId
}

func (m *KeepAliveAckMsg) Serialize(s *Serializer) {
}

func (m *KeepAliveAckMsg) Deserialize(d *Deserializer) {
}

func init() {
	registerFactory(
		KeepAliveAckMsgId,
		func() Msg { return &KeepAliveAckMsg{} },
	)
}

type ActivateCallPlaneMessage struct {
	LineInstnce uint32
}

func (m *ActivateCallPlaneMessage) Id() uint32 {
	return ActivateCallPlaneMessageId
}

func (m *ActivateCallPlaneMessage) Serialize(s *Serializer) {
	s.WriteUint32(m.LineInstnce)
}

func (m *ActivateCallPlaneMessage) Deserialize(d *Deserializer) {
	m.LineInstnce = d.ReadUint32()
}

func init() {
	registerFactory(
		ActivateCallPlaneMessageId,
		func() Msg { return &ActivateCallPlaneMessage{} },
	)
}

type RegisterRejMsg struct {
	ErrorMsg [33]byte
}

func (m *RegisterRejMsg) Id() uint32 {
	return RegisterRejMsgId
}

func (m *RegisterRejMsg) Serialize(s *Serializer) {
	s.WriteBytes(m.ErrorMsg[:])
}

func (m *RegisterRejMsg) Deserialize(d *Deserializer) {
	d.ReadBytes(m.ErrorMsg[:])
}

func init() {
	registerFactory(
		RegisterRejMsgId,
		func() Msg { return &RegisterRejMsg{} },
	)
}

type CapabilitiesReqMsg struct{}

func (m *CapabilitiesReqMsg) Id() uint32 {
	return CapabilitiesReqMsgId
}

func (m *CapabilitiesReqMsg) Serialize(s *Serializer) {
}

func (m *CapabilitiesReqMsg) Deserialize(d *Deserializer) {
}

func init() {
	registerFactory(
		CapabilitiesReqMsgId,
		func() Msg { return &CapabilitiesReqMsg{} },
	)
}

type ClearNotifyMsg struct{}

func (m *ClearNotifyMsg) Id() uint32 {
	return ClearNotifyMsgId
}

func (m *ClearNotifyMsg) Serialize(s *Serializer) {
}

func (m *ClearNotifyMsg) Deserialize(d *Deserializer) {
}

func init() {
	registerFactory(
		ClearNotifyMsgId,
		func() Msg { return &ClearNotifyMsg{} },
	)
}
