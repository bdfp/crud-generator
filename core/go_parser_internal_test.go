package core

import (
	"testing"
)

func TestFindModelName(t *testing.T) {
	modExpRep := &ModelRepresentation{ModelName: "Testing"}

	var modRep ModelRepresentation

	line := "type Testing struct {"
	if done := findModelName(&modRep, line); !done  {
		t.Error("Could not find model")
	}

	if modExpRep.ModelName != modRep.ModelName {
		t.Error("Incorrect ModelName extract", modExpRep.ModelName, modRep.ModelName)
	}

	line = "package core"
	if ok := findModelName(&modRep, line); ok  {
		t.Error("Incorrect parsing")
	}
}

func TestFindKeyName(t *testing.T) {
	modExpRep := &ModelRepresentation{
		ModelName: "Testing",
		Keys: []string{"ID"},
	}

	var modRep ModelRepresentation

	line := "ID int64 `json:\"id\"`"
	if finish := findKeyName(&modRep, line); finish  {
		t.Error("Incorrect pasing")
	}

	if modExpRep.Keys[0] != modRep.Keys[0] {
		t.Error("Incorrect ModelName extract", modExpRep.ModelName, modRep.ModelName)
	}

	line = "}"
	if finish := findKeyName(&modRep, line); !finish  {
		t.Error("Could not end parsing keys")
	}
}