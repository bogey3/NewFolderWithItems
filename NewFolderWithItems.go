package main

import (
	"os"
	"strconv"
	"strings"
)

func main(){
	files := os.Args[1:]
	baseFolderName := "New folder with items"
	folderName := baseFolderName
	index := 0
	_, err := os.Stat(folderName)
	for !os.IsNotExist(err){
		index ++
		folderName = baseFolderName + " (" + strconv.Itoa(index) + ")"
		_, err = os.Stat(folderName)
	}
	os.Mkdir(folderName, os.ModeDir)
	for _, file := range files {
		filePathSplit := strings.Split(file, string(os.PathSeparator))
		filename := filePathSplit[len(filePathSplit)-1]
		os.Rename(file, folderName + string(os.PathSeparator) + filename)

	}

}
