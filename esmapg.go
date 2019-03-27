package esmapg

import (
	"encoding/json"
)

// Map is the root of mapping
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

// Parse read config json
func Parse(config []byte) []Map {
	var hash map[string]fields
	json.Unmarshal(config, &hash)

	var maps []Map
	for name, attr := range hash {
		maps = append(maps, Map{name, attr})
	}

	return maps
}
