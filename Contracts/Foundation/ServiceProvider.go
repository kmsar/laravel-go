package Foundation

type ServiceProvider interface {

	// Register register the services.
	Register(application IApplication)

	//Boot any application services after register.
	Boot(application IApplication)

	// Start  service.
	Start() error

	// Stop  service.
	Stop()
}
