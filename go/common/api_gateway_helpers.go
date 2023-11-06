package common

import "google.golang.org/protobuf/encoding/protojson"

var JsonMarshalOptions = protojson.MarshalOptions{
	EmitUnpopulated: true,
	Multiline:       true,
}
