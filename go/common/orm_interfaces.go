package common

import "database/sql"

const (
	RECKEY = "_RK_"
)

type IORM interface {
	Introspect() IIntrospect
}

type IOrmPlugin interface {
	SQL() bool
	Write(interface{}, IORM) error
}

type SqlDatabaseDecorator interface {
	DbType() string
	Connect(...string) *sql.DB
}
