package protocol

import (
	"github.com/saichler/my.simple/go/net/model"
	"github.com/saichler/my.simple/go/types"
	"github.com/saichler/my.simple/go/utils/security"
	"google.golang.org/protobuf/proto"
	"reflect"
	"sync/atomic"
)

// Running sequence number for the messages
var sequence atomic.Int32

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

func CreateMessageFor(priority model.Priority, action model.Action, key, source, dest string, pb proto.Message) ([]byte, error) {
	//first marshal the protobuf into bytes
	data, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	//Encode the data
	encData, err := security.Encode(data, key)
	if err != nil {
		return nil, err
	}
	//create the wrapping message for the destination
	msg := &model.SecureMessage{}
	msg.Source = source
	msg.Destination = dest
	msg.Sequence = sequence.Add(1)
	msg.Priority = priority
	msg.ProtoData = encData
	msg.ProtoTypeName = reflect.ValueOf(pb).Elem().Type().Name()
	msg.Action = action
	//Now serialize the message
	msgData, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	//Create the header for the switch
	header := GenerateHeader(msg)
	//Append the msgData to the header
	header = append(header, msgData...)
	return header, nil
}
