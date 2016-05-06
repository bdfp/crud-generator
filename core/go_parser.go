package core

import (
	"bufio"
	"github.com/shakdwipeea/crudgenerator/cli"
	"log"
	"os"
	"strings"
)

//ModelRepresentation Representation of found Model
type ModelRepresentation struct {
	ModelName string
	Keys      []string
}

const (
	//FINDMODEL state of parser when finding struct definition
	FINDMODEL = iota

	//FINDKEY state of parser when findingkeys of struct
	FINDKEY = iota
)

//ModelType keyword with which Model is represented
const ModelType = "struct"

func findModelName(mod *ModelRepresentation, data string) bool {
	chunk := strings.Split(data, " ")

	for k, v := range chunk {
		if v == ModelType {
			mod.ModelName = chunk[k-1]
			return true
		}
	}

	return false

}

func findKeyName(mod *ModelRepresentation, line string) bool {
	chunk := strings.Split(line, " ")

	//keyName has the first letter capital
	keyName := chunk[0]

	if keyName != "}" {
		mod.Keys = append(mod.Keys, keyName)
		return false
	}

	return true

}

func generateModelRep(filepath *string) (*[]ModelRepresentation, error) {
	var models []ModelRepresentation
	state := FINDMODEL
	var tmpModel ModelRepresentation

	file, err := os.Open(*filepath)
	if err != nil {
		cli.ErrLog("Could not open " + *filepath)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		switch state {
		case FINDMODEL:
			if done := findModelName(&tmpModel, scanner.Text()); done {
				state = FINDKEY
			}
		case FINDKEY:
			if finish := findKeyName(&tmpModel, scanner.Text()); finish {
				state = FINDMODEL
				models = append(models, tmpModel)
			}
		}

	}

	if err = scanner.Err(); err != nil {
		cli.ErrLog(err.Error())
	}

	return &models, nil
}

//ParseFile parse the file and create a Representation that can be used by Writer
func ParseFile(filepath *string) {
	modRep, err := generateModelRep(filepath)
	if err != nil {
		cli.ErrLog(err.Error())
	}

	for _, v := range *modRep {
		cli.Log("Model found ", v.ModelName)
	}
}
