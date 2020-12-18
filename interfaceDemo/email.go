package interfaceDemo

import "fmt"

type EmailType struct {
	Name  string
	Email string
}

func (receiver *EmailType) SendEmail() {
	fmt.Println("请发送邮件给：" + receiver.Name + "，到 ：" + receiver.Email)
}
func (receiver *EmailType) ReceiveEmail() {
	fmt.Println(receiver.Name + "的邮箱：" + receiver.Email + "收到邮件！")
}
