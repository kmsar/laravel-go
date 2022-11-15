package Email

import "github.com/kmsar/laravel-go/Framework/Contracts/Support"

type EmailFactory interface {

	// Mailer Get the envelope by the given name.
	Mailer(name ...string) Mailer

	// Extend email factory with given name and mail driver.
	Extend(name string, driver MailerDriver)
}

type EmailContent interface {

	// Text Get the plain text view for the message.
	Text() string

	// Html Get the rendered HTML content for the message.
	Html() string
}

type Mailable interface {

	// SetCc Set the recipients of the message.
	SetCc(address ...string) Mailable

	// SetBcc Set cc email recipients.
	SetBcc(address ...string) Mailable

	// SetTo Set the recipients of the message.
	SetTo(address ...string) Mailable

	// SetFrom Set the from address of this message.
	SetFrom(from string) Mailable

	// Queue  the given message.
	Queue(queue string) Mailable

	// Delay Deliver the queued message after the given delay.
	Delay(delay int) Mailable

	// GetCc Get recipients of mail.
	GetCc() []string

	// GetBcc Get cc email recipients.
	GetBcc() []string

	// GetTo Get recipients of mail.
	GetTo() []string

	// GetSubject Get the subject of the message.
	GetSubject() string

	// GetFrom Get the "from" address to the message.
	GetFrom() string

	// GetText Get the plain text view for the message.
	GetText() string

	// GetHtml Get the rendered HTML content for the message.
	GetHtml() string

	// GetQueue get message queue name.
	GetQueue() string

	// GetDelay Get message queue delay.
	GetDelay() int
}

// MailerDriver Get mail driver by given name and configuration info.
type MailerDriver func(name string, config Support.Fields) Mailer

type Mailer interface {

	// Raw send a new message with only a raw text part.
	Raw(subject, text string, to []string) error

	// Send  a new message using a view.
	Send(mail Mailable) error

	// Queue queue a new e-mail message for sending.
	Queue(mail Mailable, queue ...string) error

	// Later queue a new e-mail message for sending after (n) seconds.
	Later(delay int, mail Mailable, queue ...string) error
}
