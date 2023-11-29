package common

const (
	RECKEY = "_RK_"
)

type IORM interface {
	Introspect() IIntrospect
	Persist(interface{}) error
	Fetch(IFetch) (interface{}, error)
}

type IOrmPlugin interface {
	Init(IORM, ...interface{}) error
	RelationalData() bool
	Decorator() DataStoreDecorator

	Persist(interface{}) error
	Fetch(fetch IFetch) (interface{}, error)
}

type DataStoreDecorator interface {
	DataStoreTypeName() string
	Connect(...string) interface{}
}
