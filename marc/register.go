package marc

import (
	"io"
)

var (
	decoders = make(map[string]DecoderFactory)
	encoders = make(map[string]EncoderFactory)
)

type DecoderFactory func(io.Reader) Decoder
type EncoderFactory func(io.Writer) Encoder

type Decoder interface {
	Decode() (*Record, error)
}

type Encoder interface {
	Encode(*Record) error
}

func init() {
	RegisterDecoder("marcxml", NewMARCXMLDecoder)
}

func RegisterDecoder(name string, fn DecoderFactory) {
	decoders[name] = fn
}

func RegisterEncoder(name string, fn EncoderFactory) {
	encoders[name] = fn
}

func NewDecoder(name string, r io.Reader) Decoder {
	return decoders[name](r)
}

func NewEncoder(name string, w io.Writer) Encoder {
	return encoders[name](w)
}
