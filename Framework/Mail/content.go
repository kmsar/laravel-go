package Mail

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IMail"
)

func Text(text string) IMail.EmailContent {
	return Textual{text: text}
}

type Textual struct {
	text string
}

func (t Textual) Text() string {
	return t.text
}

func (t Textual) Html() string {
	return t.text
}
