package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	files := os.Args[1:]
	pathSep := string(os.PathSeparator)
	if len(files) == 0 {
		os.Exit(1)
	}
	baseFolderName := "New folder with items"
	folderName := baseFolderName
	_, err := os.Stat(folderName)
	for index := 1; !os.IsNotExist(err); index++ {
		folderName = fmt.Sprintf("%s (%d)", baseFolderName, index)
		_,err = os.Stat(folderName)
	}
	os.Mkdir(folderName, os.ModeDir)
	for _, file := range files {
		os.Rename(file, folderName + pathSep + file[strings.LastIndex(file, pathSep)+1:])
	}
}
