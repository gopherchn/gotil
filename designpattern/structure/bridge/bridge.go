package bridge

import "fmt"

// 桥接：将抽象和实现进行解耦
// 通知和告警抽象出来
type MSGSender interface {
	Send(msg string) error
}

type Notification interface {
	Notify(msg string) error
}

type EmailMSGSender struct {
	emails []string
}

func NewEmailMSGSender(emails []string) *EmailMSGSender {
	return &EmailMSGSender{
		emails: emails,
	}
}

func (s *EmailMSGSender) Send(msg string) error {
	// send email
	fmt.Println("email send", msg)
	return nil
}


// 不同的Notification都依赖同一个Sender接口
// sender可以是EmailSender，可以是MessageSender，也可以是PhoneCall Sender
type ErrorNotification struct {
	sender MSGSender
}

func NewErrorNotification(msgSender MSGSender) *ErrorNotification {
	return &ErrorNotification{
		sender: msgSender,
	}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
