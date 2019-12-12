package ActionResult

import (
	"github.com/ugorji/go/codec"
	"net/http"
)

type MsgPack struct {
	Data interface{}
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

func (r MsgPack) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, msgpackContentType)
}

func WriteMsgPack(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, msgpackContentType)
	var mh codec.MsgpackHandle
	return codec.NewEncoder(w, &mh).Encode(obj)
}

func (r MsgPack) Render(w http.ResponseWriter) error {
	return WriteMsgPack(w, r.Data)
}
