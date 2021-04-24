package marc

type Record struct {
	Leader        Leader         `xml:"leader"`
	ControlFields []ControlField `xml:"controlfield"`
	DataFields    []DataField    `xml:"datafield"`
}

type Leader string

type ControlField struct {
	Tag   string `xml:"tag,attr"`
	Value string `xml:",chardata"`
}

type DataField struct {
	Tag       string     `xml:"tag,attr"`
	Ind1      string     `xml:"ind1,attr"`
	Ind2      string     `xml:"ind2,attr"`
	SubFields []SubField `xml:"subfield"`
}

type SubField struct {
	Code  string `xml:"code,attr"`
	Value string `xml:",chardata"`
}
