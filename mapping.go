package esmapg

import (
	"encoding/json"

	"github.com/jinzhu/inflection"
)

type doc struct {
	Properties map[string]property `json:"properties"`
}

type property struct {
	Type       string     `json:"type"`
	Properties properties `json:"properties"`
}

type properties map[string]property

// Mapping convert config json into es mapping
func (m *Map) Mapping() string {
	js, err := json.Marshal(m.Fields.mapping())
	if err != nil {
		panic(err)
	}
	return string(js)
}

func (fs *fields) mapping() properties {
	props := make(properties)

	for name, subFields := range fs.HasMany {
		props[inflection.Plural(name)] = property{"nested", subFields.mapping()}
	}

	for name, subFields := range fs.HasOne {
		props[inflection.Singular(name)] = property{"object", subFields.mapping()}
	}

	for name, subFields := range fs.BelongsTo {
		props[inflection.Singular(name)] = property{"object", subFields.mapping()}
	}

	return props
}
