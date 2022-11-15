package Mail

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IMail"
)

func ConvertToMail(mailable IMail.Mailable) *Mail {
	if mail, ok := mailable.(*Mail); ok {
		return mail
	}
	return &Mail{
		From:    mailable.GetFrom(),
		Subject: mailable.GetSubject(),
		Text:    mailable.GetText(),
		Html:    mailable.GetHtml(),
		To:      mailable.GetTo(),
		Cc:      mailable.GetCc(),
		Bcc:     mailable.GetBcc(),
		queue:   mailable.GetQueue(),
		delay:   mailable.GetDelay(),
	}
}

func New(subject string, content IMail.EmailContent) IMail.Mailable {
	return &Mail{
		Subject: subject,
		Text:    content.Text(),
		Html:    content.Html(),
		To:      make([]string, 0),
		Cc:      make([]string, 0),
		Bcc:     make([]string, 0),
	}
}

type Mail struct {
	From    string
	Subject string
	Text    string
	Html    string
	To      []string
	Cc      []string
	Bcc     []string
	queue   string
	delay   int
}

func (mail *Mail) SetCc(address ...string) IMail.Mailable {
	mail.Cc = address
	return mail
}

func (mail *Mail) SetBcc(address ...string) IMail.Mailable {
	mail.Bcc = address
	return mail
}

func (mail *Mail) SetTo(address ...string) IMail.Mailable {
	mail.To = address
	return mail
}

func (mail *Mail) Queue(queue string) IMail.Mailable {
	mail.queue = queue
	return mail
}

func (mail *Mail) Delay(delay int) IMail.Mailable {
	mail.delay = delay
	return mail
}

func (mail *Mail) GetCc() []string {
	return mail.Cc
}

func (mail *Mail) GetBcc() []string {
	return mail.Bcc
}

func (mail *Mail) GetTo() []string {
	return mail.To
}

func (mail *Mail) GetSubject() string {
	return mail.Subject
}

func (mail *Mail) SetFrom(from string) IMail.Mailable {
	mail.From = from
	return mail
}
func (mail *Mail) GetFrom() string {
	return mail.From
}

func (mail *Mail) GetText() string {
	return mail.Text
}

func (mail *Mail) GetHtml() string {
	return mail.Html
}

func (mail *Mail) GetQueue() string {
	return mail.queue
}

func (mail *Mail) GetDelay() int {
	return mail.delay
}
