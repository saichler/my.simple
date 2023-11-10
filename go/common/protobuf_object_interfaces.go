package common

type ProtobufObjectType interface {
	Add(interface{}) ([]byte, int)
	Get([]byte, int) (interface{}, int)
}
