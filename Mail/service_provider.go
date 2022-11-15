package Mail

import (
	"crypto/tls"
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IMail"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IQueue"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISerialize"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"net/smtp"
)

type ServiceProvider struct {
	app IFoundation.IApplication
}

func (service *ServiceProvider) Register(application IFoundation.IApplication) {
	service.app = application

	application.NamedSingleton("mail.factory", func(config IConfig.Config, queue IQueue.Queue) IMail.EmailFactory {
		return &Factory{
			config:  config.Get("mail").(Config),
			mailers: map[string]IMail.Mailer{},
			drivers: map[string]IMail.MailerDriver{
				"mailer": func(name string, config Support.Fields) IMail.Mailer {

					tlsConfig, _ := config["tls"].(*tls.Config)
					return &Mailer{
						name:      name,
						tlsConfig: tlsConfig,
						from:      Field.GetStringField(config, "from"),
						auth: smtp.PlainAuth(
							Field.GetStringField(config, "identity"),
							Field.GetStringField(config, "username"),
							Field.GetStringField(config, "password"),
							Field.GetStringField(config, "host"),
						),
						address: fmt.Sprintf("%s:%s", Field.GetStringField(config, "host"), Field.GetStringField(config, "port")),
						queue:   queue,
					}
				},
			},
		}
	})

	application.NamedSingleton("mailer", func(factory IMail.EmailFactory) IMail.Mailer {
		return factory.Mailer()
	})
}

func (service *ServiceProvider) Start() error {
	service.app.Call(func(serializer ISerialize.ClassSerializer) {
		serializer.Register(JobClass)
	})
	return nil
}

func (service *ServiceProvider) Stop() {
}
