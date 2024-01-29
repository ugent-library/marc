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
	RegisterDecoder("alephseq", NewAlephSeqDecoder)
	RegisterEncoder("marcinjson", NewMARCInJSONEncoder)
}

func RegisterDecoder(name string, fn DecoderFactory) {
	decoders[name] = fn
}

func RegisterEncoder(name string, fn EncoderFactory) {
	encoders[name] = fn
}

func NewDecoder(format string, r io.Reader) Decoder {
	if dec, ok := decoders[format]; ok {
		return dec(r)
	}
	return nil
}

func NewEncoder(format string, w io.Writer) Encoder {
	if enc, ok := encoders[format]; ok {
		return enc(w)
	}
	return nil
}
