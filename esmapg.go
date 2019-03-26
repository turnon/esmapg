package esmapg

import (
	"encoding/json"
	"strings"
)

type Map struct {
	Name   string
	Fields fields
}

type fields struct {
	Only      []string
	HasMany   map[string]fields `json:"has_many"`
	HasOne    map[string]fields `json:"has_one"`
	BelongsTo map[string]fields `json:"belongs_to"`
}

func (m *Map) Sql() string {
	return "SELECT " + m.onlySql() + " FROM " + m.Name
}

func (m *Map) onlySql() string {
	return strings.Join(m.Fields.Only, ", ")
}

func Parse(config []byte) []Map {
	var hash map[string]fields
	json.Unmarshal(config, &hash)

	var maps []Map
	for name, attr := range hash {
		maps = append(maps, Map{name, attr})
	}

	return maps
}
