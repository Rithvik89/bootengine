package utils

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func CreateFile(path string) {
	_, err := os.Create(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	// if its a file ending in .go add title to it
	if isGoFile(path) {
		title := "package " + getDirFromPath(path)
		os.WriteFile(path, []byte(title), 0644)
	}

	fmt.Printf("File Created at %s", path)
}

func CreateFolder(path string, root bool) {
	if root {
		_, err := git.PlainInit(path, false)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Printf("Not able to create a directory : %s", err)
		return
	}
	fmt.Printf("Directory created at %s", path)
}

func isGoFile(path string) bool {
	fileName := getFileFromPath(path)
	fileSplit := strings.Split(fileName, ".")
	if len(fileSplit) == 2 && fileSplit[1] == "go" {
		return true
	}
	return false
}

func getFileFromPath(path string) string {
	splits := strings.Split(path, "/")
	index := len(splits) - 1
	return splits[index]
}

func getDirFromPath(path string) string {
	splits := strings.Split(path, "/")
	index := len(splits) - 2
	return splits[index]
}

func RemoveDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Successfully deleted at %s", path)
}
