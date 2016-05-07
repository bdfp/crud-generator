package main

import (
	"fmt"
	"github.com/shakdwipeea/crudgenerator/cli"
	"github.com/shakdwipeea/crudgenerator/core"
	"log"
	"os"
	"strings"
)

// OUTPUT
// [proj_dir]/model/*.go are struct representation
// [proj_dir]/db/*.go are CRUD db functions
// [proj_dir]/rest/utils.go Util functions that helps in API
// [proj_dir]/rest/*.go api
//							GET /modelName
//							GET /modelName/modelId
//							POST /modelName {modelBody} ADD
// 							POST /modelName	 {modelBody} EDIT By Each param & EDIT All Param
//							DELETE /modelName/modelId DELETE

//USAGE generator [proj_dir] [modelDirName]

//WORKFLOW
// The pipeline will be sth like
// READER reads the file from the fs
//    |
// PARSER parses the file read
//    |
// WRITER writes the corres db and api handlers

//Reader reads files from the file system. Implement this to describe your file structure
type Reader interface {
	ReadDir(*string) (*[]string, error)
	ReadFile(fileName *string)
}

//Parser parses the required model files and creates the system representation
type Parser interface {
}

//Writer writes the output files in the filesystem. Implement this to describe your file structure
type Writer interface {
}

func main() {
	fmt.Println("Starting scaffolding")
	cmdArgs := os.Args[1:]

	//todo ShowBanner()

	if len(cmdArgs) != 2 {
		log.Println("Usage is generator [proj_dir] [modelDirName]")
		os.Exit(1)
	}

	projectDir := cmdArgs[0]
	modelDirName := cmdArgs[1]

	cwd, err := os.Getwd()
	if err != nil {
		panic("Error Reading CWD")
	}

	cli.Log("Current Directory is", cwd)

	wModelDir := cwd + "/" + projectDir + modelDirName

	cli.Log("Reading Models from ", wModelDir)

	c := strings.Split(projectDir, "/")
	projName := c[len(c) - 1]

	reader := &core.GolangReader{
		ProjectDir: cwd + "/" + projectDir,
		ProjectName: projName,
	}
	filePaths, err := reader.ReadDir(&wModelDir)
	if err != nil {
		panic("Could not read models " + err.Error())
	}

	for _, v := range *filePaths {
		reader.ParseFile(&v)
		reader.WriteFiles(&v)
	}

}

func printArr(arr *[]string) {
	for k, v := range *arr {
		cli.Log("Key is "+string(k), " Value is "+string(v))
	}
}
