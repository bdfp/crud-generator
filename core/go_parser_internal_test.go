package core

import (
	"testing"
)

func TestFindModelName(t *testing.T) {
	modExpRep := &ModelRepresentation{Name: "Testing"}

	var modRep ModelRepresentation

	line := "type Testing struct {"
	if done := findModelName(&modRep, line); !done  {
		t.Error("Could not find model")
	}

	if modExpRep.Name != modRep.Name {
		t.Error("Incorrect ModelName extract", modExpRep.Name, modRep.Name)
	}

	line = "package core"
	if ok := findModelName(&modRep, line); ok  {
		t.Error("Incorrect parsing")
	}
}

func TestFindKeyName(t *testing.T) {
	modExpRep := &ModelRepresentation{
		Name: "Testing",
		Keys: []string{"ID"},
	}

	var modRep ModelRepresentation

	line := "ID int64 `json:\"id\"`"
	if finish := findKeyName(&modRep, line); finish  {
		t.Error("Incorrect pasing")
	}

	if modExpRep.Keys[0] != modRep.Keys[0] {
		t.Error("Incorrect ModelName extract", modExpRep.Name, modRep.Name)
	}

	line = "}"
	if finish := findKeyName(&modRep, line); !finish  {
		t.Error("Could not end parsing keys")
	}
}