package service

import (
)

type DataLayer interface {
	Ping() (string, error)
}
