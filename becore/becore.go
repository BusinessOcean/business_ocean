package becore

type BeCore interface {

	// IsDev returns whether the app is in dev mode.
	IsDev() bool

	// Bootstrap takes care for initializing the application
	// (open db connections, load settings, config setting, etc.).
	//
	// It will call ResetBootstrapState() if the application was already bootstrapped.
	Bootstrap() error

	// ResetBootstrapState takes care for releasing initialized app resources
	// (eg. closing db connections).
	ResetBootstrapState() error
}
