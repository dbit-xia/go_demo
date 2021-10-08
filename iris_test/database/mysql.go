package database

import "fmt"

type mysql struct{}

func (db *mysql) Exec(q string) error {
	return fmt.Errorf("mysql: not implemented <%s>", q)
}
