package common

type IORM interface {
	Introspect() IIntrospect
}

type IOrmPlugin interface {
	SQL() bool
	Write(interface{}, IORM) error
}
