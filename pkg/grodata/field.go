package grodata

import (
	"emperror.dev/errors"
	"encoding/json"
	"strings"
)

type Fields []Field

func (f Fields) GetFields() []string {
	var fields []string
	for _, field := range f {
		fields = append(fields, field.TypeName)
	}
	return fields
}

func (f Fields) GetField(name string) (Field, bool) {
	for _, field := range f {
		if field.TypeName == name {
			return field, true
		}
	}
	return Field{}, false
}

type StringOrFields struct {
	Type   string
	Str    []string           `json:"string,omitempty"`
	Fields []map[string]Field `json:"fields,omitempty"`
}

func (s *StringOrFields) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		s.Type = "primitive"
		s.Str = []string{str}
		s.Fields = nil
		return nil
	}
	var strs []string
	if err := json.Unmarshal(data, &strs); err == nil {
		s.Type = "controlledVocabulary"
		s.Str = strs
		s.Fields = nil
		return nil
	}

	var fields = []map[string]Field{}
	if err := json.Unmarshal(data, &fields); err == nil {
		s.Type = "compound"
		s.Str = nil
		s.Fields = fields
		return nil
	} else {
		return errors.Wrapf(err, "error unmarshalling StringOrFields '%s'", string(data))
	}
}

func (s StringOrFields) MarshalJSON() ([]byte, error) {
	if s.String != nil {
		return json.Marshal(s.String)
	}
	return json.Marshal(s.Fields)
}

func (s StringOrFields) GetField(name string) (Field, bool) {
	for _, field := range s.Fields {
		for k, f := range field {
			if k == name {
				return f, true
			}
		}
	}
	return Field{}, false
}

func (s StringOrFields) String() string {
	if s.Str != nil {
		return strings.Join(s.Str, "; ")
	}
	var str string
	for _, field := range s.Fields {
		for k, f := range field {
			str += k + ": " + f.Value.String() + " // "
		}
	}
	return str
}

type Field struct {
	TypeName  string         `json:"typeName,omitempty"`
	Multiple  bool           `json:"multiple,omitempty"`
	TypeClass string         `json:"typeClass,omitempty"`
	Value     StringOrFields `json:"value,omitempty"`
}
