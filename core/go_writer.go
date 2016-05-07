package core

import (
	"text/template"
	"io/ioutil"
	"os"
	"strings"
)

type DbHandlerData struct {
	ModelData string
	modelData string
	Keys []string
}

//WriteApiFiles Write handlers to the file system
func WriteApiFiles(modelName *ModelRepresentation, projDir string) error {
	targetApiFile := projDir + "/" + modelName.Name + ".go"
	f, err := os.Create(targetApiFile)
	if err != nil {
		return err
	}

	t := template.New("ApiHandler template")

	//read template input
	dat, err := ioutil.ReadFile("../template/apiHandler.ego")
	if err != nil {
		return err
	}

	t, err = t.Parse(string(dat))
	if err != nil {
		return err
	}

	dat = &DbHandlerData{
		ModelData: modelName.Name,
		modelData: strings.ToLower(modelName.Name[:1]) + modelName.Name[1:],
	}

	return t.Execute(f, dat)
}

func (r *GolangReader) WriteFiles(filepath *string, modelRep []ModelRepresentation) {

	for _, v := range modelRep {
		WriteApiFiles(&v, r.ProjectDir)
	}
}
