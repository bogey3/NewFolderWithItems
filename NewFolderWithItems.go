package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	files := os.Args[1:]
	baseFolderName := "New folder with items"
	folderName := baseFolderName
	_, err := os.Stat(folderName)
	for index := 1; !os.IsNotExist(err); index++ {
		folderName = fmt.Sprintf("%s (%d)", baseFolderName, index)
		_,err = os.Stat(folderName)
	}
	os.Mkdir(folderName, os.ModeDir)
	for _, file := range files {
		filePathSplit := strings.Split(file, string(os.PathSeparator))
		filename := filePathSplit[len(filePathSplit)-1]
		os.Rename(file, folderName + string(os.PathSeparator) + filename)

	}
}
