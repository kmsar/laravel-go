package signal

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"os"
	"os/signal"
	"syscall"
)

type ServiceProvider struct {
	signals       []os.Signal
	signalChannel chan os.Signal
	app           IFoundation.IApplication
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application
}

func (this *ServiceProvider) Start() (err error) {
	this.signalChannel = make(chan os.Signal)
	signal.Notify(this.signalChannel, this.signals...)
	for sign := range this.signalChannel {
		switch sign {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			this.app.Stop()
		}
	}

	return err
}
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}
func (this *ServiceProvider) Stop() {
	close(this.signalChannel)
}
