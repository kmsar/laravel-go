package IMail

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type EmailFactory interface {
	//Mailer gets the mail driver by the given name
	// Get the envelope by the given name.
	Mailer(name ...string) Mailer

	//Extend extends the email factory by the given name and email driver
	// Extend email factory with given name and mail driver.
	Extend(name string, driver MailerDriver)
}

type EmailContent interface {
	//Text Gets the plain text view of the message
	// Get the plain text view for the message.
	Text() string

	// Html gets the rendered HTML content of the message
	// Get the rendered HTML content for the message.
	Html() string
}

type Mailable interface {
	// SetCc set the recipient of the mail
	// Set the recipients of the message.
	SetCc(address ...string) Mailable

	//SetBcc set email recipients who are CC'd
	// Set cc email recipients.
	SetBcc(address ...string) Mailable

	// SetTo Set the recipient of the mail
	// Set the recipients of the message.
	SetTo(address ...string) Mailable

	// SetFrom 设置此消息的发件人地址
	// Set the from address of this message.
	SetFrom(from string) Mailable

	// Queue 将给定的消息加入队列
	// queue the given message.
	Queue(queue string) Mailable

	// Delay 在给定的延迟后传递队列的消息
	// Deliver the queued message after the given delay.
	Delay(delay int) Mailable

	// GetCc 获取邮件的收件人
	// Get recipients of mail.
	GetCc() []string

	// GetBcc 获取抄送邮件收件人
	// Get cc email recipients.
	GetBcc() []string

	// GetTo 获取邮件的收件人
	// Get recipients of mail.
	GetTo() []string

	// GetSubject 获取消息的主题
	// Get the subject of the message.
	GetSubject() string

	// GetFrom 获取消息的发件人地址
	// Get the "from" address to the message.
	GetFrom() string

	// GetText 获取消息的纯文本视图
	// Get the plain text view for the message.
	GetText() string

	// GetHtml 获取消息的呈现 HTML 内容
	// Get the rendered HTML content for the message.
	GetHtml() string

	// GetQueue 获取消息队列名称
	// get message queue name.
	GetQueue() string

	// GetDelay 获取消息队列延迟
	// Get message queue delay.
	GetDelay() int
}

// MailerDriver gets the mail driver with the given name and configuration information
// Get mail driver by given name and configuration info.
type MailerDriver func(name string, data Support.Fields) Mailer

type Mailer interface {
	// Raw 发送仅包含原始文本部分的新消息
	// send a new message with only a raw text part.
	Raw(subject, text string, to []string) error

	// Send 使用视图发送新消息
	// send a new message using a view.
	Send(mail Mailable) error

	// Queue 排队等待发送的新电子邮件
	// queue a new e-mail message for sending.
	Queue(mail Mailable, queue ...string) error

	// Later 在 (n) 秒后排队发送新的电子邮件消息
	// queue a new e-mail message for sending after (n) seconds.
	Later(delay int, mail Mailable, queue ...string) error
}
