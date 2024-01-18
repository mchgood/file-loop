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

	err = ioutil.WriteFile("out.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output saved to out.json")
}

func ListFiles(dirPath string) ([]FileInfo, error) {
	var fileList []FileInfo

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file := FileInfo{
				Path: path,
				Name: info.Name(),
			}
			fileList = append(fileList, file)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
