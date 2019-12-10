package ActionResult

import (
	"github.com/golang/protobuf/proto"
	"net/http"
)

type ProtoBuf struct {
	Data interface{}
}

var protobufContentType = []string{"application/x-protobuf"}

func (r ProtoBuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	bytes, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

// WriteContentType (ProtoBuf) writes ProtoBuf ContentType.
func (r ProtoBuf) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, protobufContentType)
}
