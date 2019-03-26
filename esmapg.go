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
	return "SELECT " + m.Fields.sql(m.Name) + " FROM " + m.Name
}

func (fs *fields) sql(parentName string) string {
	allSql := []string{fs.onlySql(), fs.belongsToSql(parentName)}

	var sqls []string
	for _, sql := range allSql {
		if sql != "" {
			sqls = append(sqls, sql)
		}
	}

	return strings.Join(sqls, ", ")
}

func (fs *fields) onlySql() string {
	return strings.Join(fs.Only, ", ")
}

func (fs *fields) belongsToSql(parentName string) string {
	var sqls []string
	for name, subFields := range fs.BelongsTo {
		joining := parentName + "." + name + "_id = " + name + ".id"
		sql := "( SELECT row_to_json(t) FROM ( SELECT " +
			subFields.sql(name) + " FROM " + name + " WHERE " + joining +
			") t ) AS " + name
		sqls = append(sqls, sql)
	}
	return strings.Join(sqls, ", ")
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
