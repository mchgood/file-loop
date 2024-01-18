package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func main() {
	files, err := ListFiles(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonData, err := json.MarshalIndent(files, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = WriteFileWithUTF8("out.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output saved to out.json")
}

func ListFiles(dirPath string) ([]FileInfo, error) {
	var fileList []FileInfo

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileInfo := FileInfo{
				Path: filepath.Join(dirPath, file.Name()),
				Name: file.Name(),
			}
			fileList = append(fileList, fileInfo)
		}
	}

	return fileList, nil
}

func WriteFileWithUTF8(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
