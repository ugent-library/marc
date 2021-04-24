package marc

import (
	"encoding/xml"
	"io"
)

type marcxmlDecoder struct {
	xmlDecoder *xml.Decoder
}

func NewMARCXMLDecoder(r io.Reader) Decoder {
	return &marcxmlDecoder{xml.NewDecoder(r)}
}

func (d *marcxmlDecoder) Decode() (*Record, error) {
	for {
		dec := d.xmlDecoder
		tok, err := dec.Token()
		if tok == nil || err == io.EOF {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		switch ty := tok.(type) {
		case xml.StartElement:
			if ty.Name.Local == "record" {
				rec := Record{}
				if err = dec.DecodeElement(&rec, &ty); err != nil {
					return nil, err
				}
				return &rec, nil
			}
		}
	}
}
