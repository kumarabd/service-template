package service

import (
	"github.com/realnighthawk/bucky/errors"
)

var (
	ErrKubeNewCode = "test"
)

func ErrKubeNew(err error) error {
	return errors.New(ErrKubeNewCode, errors.Alert, "Error parsing message", err.Error())
}
