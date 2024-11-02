package corron

import (
	"encoding/json"
	"io"
)

type SelfConscious interface {
	Validate() error
}

func Unmarshal(from []byte, to SelfConscious) error {
	if err := json.Unmarshal(from, to); err != nil {
		return err
	}

	return to.Validate()
}

type Decoder struct {
	*json.Decoder
}

func NewDecoder(r io.Reader) Decoder {
	return Decoder{
		Decoder: json.NewDecoder(r),
	}
}

func (decoder Decoder) Decode(to SelfConscious) error {
	if err := decoder.Decoder.Decode(to); err != nil {
		return err
	}

	return to.Validate()
}
