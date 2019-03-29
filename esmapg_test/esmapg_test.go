package esmapg_test

import (
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/turnon/esmapg"
)

var (
	reRmSpaces = regexp.MustCompile(`\s|\n`)
	maps       = esmapg.Parse(jsonData())
)

func jsonData() []byte {
	json, _ := ioutil.ReadFile("./config.json")
	return json
}

func sqlData() string {
	sql, _ := ioutil.ReadFile("./select.sql")
	return rmSpaces(string(sql))
}

func mappingsData() string {
	mappings, _ := ioutil.ReadFile("./mappings.json")
	return rmSpaces(string(mappings))
}

func rmSpaces(str string) string {
	return reRmSpaces.ReplaceAllString(str, ``)
}

func TestSQL(t *testing.T) {
	for _, m := range maps {
		sql := m.SQL()
		t.Log(sql)
		if rmSpaces(sql) != sqlData() {
			t.Error("wrong sql")
		}
	}
}

func TestMappings(t *testing.T) {
	for _, m := range maps {
		mappings := m.Mappings()
		t.Log(mappings)
		if rmSpaces(mappings) != mappingsData() {
			t.Error("wrong mappings")
		}
	}
}

func TestPath(t *testing.T) {
	for _, m := range maps {
		m.Paths()
		// t.Log(mappings)
		// if rmSpaces(mappings) != mappingsData() {
		// 	t.Error("wrong mappings")
		// }
	}
}
