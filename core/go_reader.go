package core

import (
	"github.com/shakdwipeea/crudgenerator/cli"
	"io/ioutil"
)

//GolangReader Reader for Golang
type GolangReader struct{
	ProjectDir string
	ProjectName string
}

//ReadDir get all filenames of a directory
func (r *GolangReader) ReadDir(dirName *string) (*[]string, error) {
	var filePaths []string

	files, err := ioutil.ReadDir(*dirName)
	if err != nil {
		return &filePaths, err
	}

	for _, v := range files {
		cli.Log("Found model ", v.Name())
		filePaths = append(filePaths, *dirName+"/"+v.Name())
	}

	return &filePaths, nil
}
