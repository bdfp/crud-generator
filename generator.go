package main

import (
	"fmt"
	"os"
)

// OUTPUT
// [proj_dir]/mdoel/*.go are struct representation
// [proj_dir]/db/*.go are CRUD db functions
// [proj_dir]/rest/utils.go Util functions that helps in API
// [proj_dir]/rest/*.go api
//							GET /modelName
//							GET /modelName/modelId
//							POST /modelName {modelBody} ADD
// 							POST /modelName	 {modelBody} EDIT By Each param & EDIT All Param
//							DELETE /modelName/modelId DELETE

//USAGE generator [proj_dir]

//WORKFLOW
// The pipeline will be sth like
// READER reads the file from the fs
//    |
// PARSER parses the file read
//    |
// WRITER writes the corres db and api handlers

//Reader reads files from the file system. Implement this to describe your file structure
type Reader interface {
}

//Parser parses the required model files and creates the system representation
type Parser interface {
}

//Writer writes the output files in the filesystem. Implement this to describe your filestructure
type Writer interface {
}

func main() {
	fmt.Println("Starting scaffolding")
	cmdArgs := os.Args[1:]

}
