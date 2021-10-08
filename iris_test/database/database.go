package database

import "app/environment"

type DB interface {
	Exec(q string) error
}

func NewDB(env environment.Env) DB {
	switch env {
	case environment.PROD:
		return &mysql{}
	case environment.DEV:
		return &sqlite{}
	default:
		panic("unknown environment")
	}
}
