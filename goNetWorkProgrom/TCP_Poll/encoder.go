package poll

import (
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}


// func (e *Encoder) Encode(message string) error {
// 	bytes.Buffer
// }
