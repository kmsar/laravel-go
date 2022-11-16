package Mail

import (
	"crypto/tls"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IMail"
	"github.com/kmsar/laravel-go/Framework/Contracts/IQueue"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Field"

	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
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
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
