package extension

import (
	"encoding/json"
	"io"
)

type DefaultJsonEncoder struct {
}

func (jsonEncoder DefaultJsonEncoder) Encode(w io.Writer, data interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(&data)
}
