package Mail

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IMail"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

type Factory struct {
	config  Config
	mailers map[string]IMail.Mailer
	drivers map[string]IMail.MailerDriver
}

func (factory *Factory) Mailer(name ...string) IMail.Mailer {
	mailer := factory.config.Default
	if len(name) > 0 {
		mailer = name[0]
	}

	return factory.getMailer(mailer)
}

func (factory *Factory) getMailer(name string) IMail.Mailer {
	if factory.mailers[name] == nil {
		config := factory.config.Mailers[name]
		if config == nil {
			panic(Exception{Exception: Exceptions.New("factory.getMailer: mailer does not exist", config)})
		}

		if driver, ok := factory.drivers[Field.GetStringField(config, "driver")]; ok {
			factory.mailers[name] = driver(name, config)
		} else {
			panic(Exception{Exception: Exceptions.New("factory.getMailer: driver does not exist", config)})
		}
	}

	return factory.mailers[name]
}

func (factory *Factory) Extend(name string, driver IMail.MailerDriver) {
	factory.drivers[name] = driver
}
