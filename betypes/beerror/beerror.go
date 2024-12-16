package beerror

import (
	"beconst"
	"fmt"
)

type BeError struct {
	key     string
	title   string
	message string
	errStr  string
}

func NewBeError(key, title, message string) *BeError {
	return &BeError{
		key:     key,
		title:   title,
		message: message,
		errStr:  fmt.Sprintf(`{"key": "%s", "title": "%s", "message": "%s"}`, key, title, message),
	}
}

func (e *BeError) Error() string {
	return e.errStr
}

func (e *BeError) Wrap(cause error) *BeError {
	e.errStr = fmt.Sprintf(`%s | Cause: %s`, e.errStr, cause.Error())
	return e
}

func (e *BeError) Is(target error) bool {
	t, ok := target.(*BeError)
	return ok && t.key == e.key
}

var (
	ErrSetupDomain        = NewBeError(beconst.ErrInDomainSetup, "Setup Domain Error", "Error setting up domain")
	ErrRegisterNilHandler = NewBeError(beconst.ErrRegisterNilHandler, "Register Nil Handler Error", "Error registering nil handler")
)
