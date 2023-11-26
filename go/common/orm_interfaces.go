package common

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
