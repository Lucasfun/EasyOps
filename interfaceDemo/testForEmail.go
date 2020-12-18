package interfaceDemo

type SendReceiver interface { //EmailType中确实实现且被本包需要的方法，本interface可代替EmailType/*EmailType
	SendEmail()
	ReceiveEmail()
}

func TestForEmail(receiver SendReceiver) {
	receiver.SendEmail()
	receiver.ReceiveEmail()
}
