package sdk

import "errors"

var (
	ErrInvalidPackageName = errors.New("invalid package name")
	ErrInvalidExtName     = errors.New("invalid extension name")
)
