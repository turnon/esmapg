package esmapg_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/turnon/esmapg"
)

func jsonData() []byte {
	json, _ := ioutil.ReadFile("./config.json")
	return json
}

func TestNew(t *testing.T) {
	maps := esmapg.Parse(jsonData())
	// fmt.Println(maps)

	for _, m := range maps {
		fmt.Println(m.SQL())
	}
}
