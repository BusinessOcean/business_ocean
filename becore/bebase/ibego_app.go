package bebase

type IBootstrap interface {
	// RefreshSettings() error
	// IsBootstrapped() bool
	Bootstrap() error
	// ResetBootstrapState() error
	// Restart() error
	// OnBeforeBootstrap() //*hook.Hook[*BootstrapEvent]
	// OnAfterBootstrap()  //*hook.Hook[*BootstrapEvent]
	// OnBeforeServe()     //*hook.Hook[*ServeEvent]
	// OnBeforeApiError()  //*hook.Hook[*ApiErrorEvent]
	// OnAfterApiError()   //*hook.Hook[*ApiErrorEvent]
	OnTerminate() error //*hook.Hook[*TerminateEvent]
}

type IConfigs interface {
	IsDev() bool
}

type ISettingAndBackup interface {
	// NewFilesystem() (*filesystem.System, error)
	// NewBackupsFilesystem() (*filesystem.System, error)
	// CreateBackup(ctx context.Context, name string) error
	// RestoreBackup(ctx context.Context, name string) error
	// OnSettingsListRequest() *hook.Hook[*SettingsListEvent]
	// OnSettingsBeforeUpdateRequest() *hook.Hook[*SettingsUpdateEvent]
	// OnSettingsAfterUpdateRequest() *hook.Hook[*SettingsUpdateEvent]

}

type IAdmin interface {
	// OnAdminsListRequest()                      //*hook.Hook[*AdminsListEvent]
	// OnAdminViewRequest()                       //*hook.Hook[*AdminViewEvent]
	// OnAdminBeforeCreateRequest()               //*hook.Hook[*AdminCreateEvent]
	// OnAdminAfterCreateRequest()                //*hook.Hook[*AdminCreateEvent]
	// OnAdminBeforeUpdateRequest()               //*hook.Hook[*AdminUpdateEvent]
	// OnAdminAfterUpdateRequest()                //*hook.Hook[*AdminUpdateEvent]
	// OnAdminBeforeDeleteRequest()               //*hook.Hook[*AdminDeleteEvent]
	// OnAdminAfterDeleteRequest()                //*hook.Hook[*AdminDeleteEvent]
	// OnAdminAuthRequest()                       //*hook.Hook[*AdminAuthEvent]
	// OnAdminBeforeAuthWithPasswordRequest()     //*hook.Hook[*AdminAuthWithPasswordEvent]
	// OnAdminAfterAuthWithPasswordRequest()      //*hook.Hook[*AdminAuthWithPasswordEvent]
	// OnAdminBeforeAuthRefreshRequest()          //*hook.Hook[*AdminAuthRefreshEvent]
	// OnAdminAfterAuthRefreshRequest()           //*hook.Hook[*AdminAuthRefreshEvent]
	// OnAdminBeforeRequestPasswordResetRequest() //*hook.Hook[*AdminRequestPasswordResetEvent]
	// OnAdminAfterRequestPasswordResetRequest()  //*hook.Hook[*AdminRequestPasswordResetEvent]
	// OnAdminBeforeConfirmPasswordResetRequest() //*hook.Hook[*AdminConfirmPasswordResetEvent]
	// OnAdminAfterConfirmPasswordResetRequest()  //*hook.Hook[*AdminConfirmPasswordResetEvent]

}

type IRealtime interface {
	// OnRealtimeConnectRequest() *hook.Hook[*RealtimeConnectEvent]
	// OnRealtimeDisconnectRequest() *hook.Hook[*RealtimeDisconnectEvent]
	// OnRealtimeBeforeMessageSend() *hook.Hook[*RealtimeMessageEvent]
	// OnRealtimeAfterMessageSend() *hook.Hook[*RealtimeMessageEvent]
	// OnRealtimeBeforeSubscribeRequest() *hook.Hook[*RealtimeSubscribeEvent]
	// OnRealtimeAfterSubscribeRequest() *hook.Hook[*RealtimeSubscribeEvent]

}
