package extension

import "io"

type Encoder interface {
	Encode(w io.Writer, data interface{}) error
}
