package marc

import (
	"encoding/json"
	"io"
)

type mijRecord struct {
	Leader string                   `json:"leader"`
	Fields []map[string]interface{} `json:"fields"`
}

type mijData struct {
	SubFields []map[string]string `json:"subfields"`
	Ind1      string              `json:"ind1"`
	Ind2      string              `json:"ind2"`
}

type mijEncoder struct {
	jsonEncoder *json.Encoder
}

func NewMARCInJSONEncoder(w io.Writer) Encoder {
	return &mijEncoder{json.NewEncoder(w)}
}

func (d *mijEncoder) Encode(r *Record) error {
	nFields := len(r.ControlFields) + len(r.DataFields)
	mijr := mijRecord{Leader: r.Leader, Fields: make([]map[string]interface{}, 0, nFields)}
	for _, f := range r.ControlFields {
		mijr.Fields = append(mijr.Fields, map[string]interface{}{f.Tag: f.Value})
	}
	for _, f := range r.DataFields {
		mijd := mijData{
			Ind1:      f.Ind1,
			Ind2:      f.Ind2,
			SubFields: make([]map[string]string, 0, len(f.SubFields)),
		}
		for _, sf := range f.SubFields {
			mijd.SubFields = append(mijd.SubFields, map[string]string{sf.Code: sf.Value})
		}
		mijr.Fields = append(mijr.Fields, map[string]interface{}{f.Tag: mijd})
	}
	return d.jsonEncoder.Encode(&mijr)
}
