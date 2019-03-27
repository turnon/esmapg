package esmapg

import (
	"encoding/json"
	"strings"

	"github.com/jinzhu/inflection"
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

// SQL convert config json into sql
func (m *Map) SQL() string {
	parentTable := inflection.Plural(m.Name)
	return "SELECT " + m.Fields.sql(parentTable) + " FROM " + parentTable
}

func (fs *fields) sql(parentTable string) string {
	allSQL := []string{fs.onlySQL(), fs.belongsToSQL(parentTable), fs.hasOneSQL(parentTable)}

	var sqls []string
	for _, sql := range allSQL {
		if sql != "" {
			sqls = append(sqls, sql)
		}
	}

	return strings.Join(sqls, ", ")
}

func (fs *fields) onlySQL() string {
	return strings.Join(fs.Only, ", ")
}

func (fs *fields) belongsToSQL(parentTable string) string {
	var sqls []string
	for childName, subFields := range fs.BelongsTo {
		childName = inflection.Singular(childName)
		childTable := inflection.Plural(childName)
		joining := parentTable + "." + childName + "_id = " + childTable + ".id"
		sql := "( SELECT row_to_json(t) FROM ( SELECT " +
			subFields.sql(childTable) + " FROM " + childTable + " WHERE " + joining +
			") t ) AS " + childName
		sqls = append(sqls, sql)
	}
	return strings.Join(sqls, ", ")
}

func (fs *fields) hasOneSQL(parentTable string) string {
	var sqls []string
	for childName, subFields := range fs.HasOne {
		parentdName := inflection.Singular(parentTable)
		childTable := inflection.Plural(childName)
		joining := parentTable + ".id = " + childTable + "." + parentdName + "_id"
		sql := "( SELECT row_to_json(t) FROM ( SELECT " +
			subFields.sql(childTable) + " FROM " + childTable + " WHERE " + joining +
			") t ) AS " + childName
		sqls = append(sqls, sql)
	}
	return strings.Join(sqls, ", ")
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
