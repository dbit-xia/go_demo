package utils

import (
	pkgErrors "github.com/dbit-xia/errors"
)

func WrapError(err error) error {
	return pkgErrors.WithStack(err,1)
}

//限制第三方包的使用范围
type Errors struct{
	Skip int
}

func (e *Errors) New(text string) error {
	return pkgErrors.New(text, e.Skip)
}
func (e *Errors) UnWrap(err error) error {
	return pkgErrors.Unwrap(err)
}
func (e *Errors) Is(err error, target error) bool {
	return pkgErrors.Is(err, target)
}
func (e *Errors) As(err error, target interface{}) bool {
	return pkgErrors.As(err, target)
}

func (e *Errors) WithStack(err error) error {
	return pkgErrors.WithStack(err, e.Skip)
}

func (e *Errors) Wrap(err error, message string) error {
	return pkgErrors.Wrap(err, message, e.Skip)
}
