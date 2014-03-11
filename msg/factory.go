package msg

var registry = map[uint32]func() Msg{}

func registerFactory(msgId uint32, factory func() Msg) {
	if _, ok := registry[msgId]; ok {
		panic(msgId)
	}

	registry[msgId] = factory
}

func newFromId(msgId uint32) Msg {
	var m Msg

	if factory, ok := registry[msgId]; ok {
		m = factory()
	} else {
		m = &GenericMsg{MsgId: msgId}
	}

	return m
}
