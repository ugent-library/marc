package marc

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var fieldRe = regexp.MustCompile("^([0-9]{9}) ([A-Z0-9]{3})(.)(.) L (.*)")

type alephseqDecoder struct {
	scanner  *bufio.Scanner
	fieldBuf []string
}

func NewAlephSeqDecoder(r io.Reader) Decoder {
	return &alephseqDecoder{scanner: bufio.NewScanner(r)}
}

func (d *alephseqDecoder) Decode() (*Record, error) {
	var id string
	var rec *Record

	if d.fieldBuf != nil {
		rec = addField(rec, d.fieldBuf)
	}

	scanner := d.scanner
	for scanner.Scan() {
		field := fieldRe.FindStringSubmatch(scanner.Text())
		if field == nil {
			continue
		}
		if id == "" {
			id = field[1]
		} else if field[1] != id {
			d.fieldBuf = field
			break
		}

		rec = addField(rec, field)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return rec, nil
}

func addField(rec *Record, field []string) *Record {
	if rec == nil {
		rec = &Record{}
	}

	tag := field[2]
	ind1 := field[3]
	ind2 := field[4]
	val := field[5]

	if tag == "LDR" {
		rec.Leader = strings.ReplaceAll(val, "^", " ")
	} else if tag == "FMT" || strings.HasPrefix(tag, "00") {
		rec.ControlFields = append(rec.ControlFields, ControlField{Tag: tag, Value: strings.ReplaceAll(val, "^", " ")})
	} else {
		f := DataField{Tag: tag, Ind1: ind1, Ind2: ind2}
		for _, v := range strings.Split(val, "$$") {
			if len(v) == 0 {
				continue
			}
			runes := []rune(v)
			f.SubFields = append(f.SubFields, SubField{Code: string(runes[0]), Value: string(runes[1:])})
		}
		rec.DataFields = append(rec.DataFields, f)
	}

	return rec
}
