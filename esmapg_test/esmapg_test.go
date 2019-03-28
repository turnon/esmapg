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

func rmSpaces(str string) string {
	return reRmSpaces.ReplaceAllString(str, ``)
}

func TestSQL(t *testing.T) {
	for _, m := range maps {
		t.Log(m.SQL())
		if rmSpaces(m.SQL()) != sqlData() {
			t.Error("wrong sql")
		}
	}
}

func TestMapping(t *testing.T) {
	for _, m := range maps {
		t.Log(m.Mapping())
	}
}
