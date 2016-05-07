package core_test

import (
	"os"
	"testing"
	"github.com/shakdwipeea/crudgenerator/core"
	"reflect"
)

type TestModel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Feature string `json:"feature"`
}

type TestSecondModel struct {
	ID      int    `json:"id"`
	Status    string `json:"status"`
}

func TestParseFile(t *testing.T) {
	curDir, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	curFilePath := curDir + "/" + "go_parser_test.go"

	r := &core.GolangReader{}
	modRep, err := r.ParseFile(&curFilePath)

	if len(modRep) != 2 {
		t.Error("All models not found")
	}

	if modRep[0].Name != "TestModel" &&
	reflect.DeepEqual(modRep[0].Keys, []string{"ID", "Name", "Feature"}) {
		t.Error("First model identified incorrectly")
	}

	if modRep[1].Name != "TestSecondModel" &&
	reflect.DeepEqual(modRep[1].Keys, []string{"ID", "Status"}) {
		t.Error("Second model identified incorrectly")
	}

}
