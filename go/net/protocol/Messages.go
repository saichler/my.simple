package protocol

import (
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/types"
	"github.com/saichler/my.simple/go/utils/security"
	"google.golang.org/protobuf/proto"
)

func GenerateHeader(msg *model.SecureMessage) []byte {
	header := make([]byte, 73)
	for i, c := range msg.Source {
		header[i] = byte(c)
	}
	for i, c := range msg.Destination {
		header[i+36] = byte(c)
	}
	header[72] = byte(msg.Priority)
	return header
}

func HeaderOf(data []byte) (string, string, model.Priority) {
	//Source will always be Uuid
	source := string(data[0:36])
	//Destination, either than being a uuid can also be a topic/multicast so it might not be full 16 bytes
	dest := make([]byte, 0)
	for i := 36; i < 72; i++ {
		if data[i] == 0 {
			break
		}
		dest = append(dest, data[i])
	}
	pri := model.Priority(data[72])
	return source, string(dest), pri
}

func MessageOf(data []byte) (*model.SecureMessage, error) {
	msg := &model.SecureMessage{}
	err := proto.Unmarshal(data[73:], msg)
	return msg, err
}

func ProtoOf(msg *model.SecureMessage, key string) (proto.Message, error) {
	data, err := security.Decode(msg.ProtoData, key)
	if err != nil {
		return nil, err
	}

	pbi, err := types.Types.New(msg.ProtoTypeName)
	if err != nil {
		return nil, err
	}

	pb := pbi.(proto.Message)
	err = proto.Unmarshal(data, pb)
	return pb, err
}