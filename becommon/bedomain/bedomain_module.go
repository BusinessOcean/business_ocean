package bedomain

import (
	"becommon/beevent"
	"becore/behook"
)

type IBeDomainModule interface {

	// RefreshSettings() error
	// IsBootstrapped() bool
	Bootstrap() error

	Run() error
	// ResetBootstrapState() error
	// Restart() error
	OnBeforeBootstrap() *behook.Hook[*beevent.ExampleHookEvent]
	// OnAfterBootstrap() *behook.Hook[*ExampleHookEvent]
	// OnBeforeServe() *behook.Hook[*ExampleHookEvent]
	// OnBeforeApiError() *behook.Hook[*ExampleHookEvent]
	// OnAfterApiError() *behook.Hook[*ExampleHookEvent]
	OnTerminate() *behook.Hook[*beevent.ExampleHookEvent]
}
