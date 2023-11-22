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
		title := "package " + getDirFromFilePath(path)
		os.WriteFile(path, []byte(title), 0644)
	}

	fmt.Printf("File Created at %s", path)
}

func CreateFolder(path string, root bool) *git.Repository {
	if root {
		repo, err := git.PlainInit(path, false)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return repo
	}
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Printf("Not able to create a directory : %s", err)
		return nil
	}
	fmt.Printf("Directory created at %s", path)
	return nil
}

func isGoFile(path string) bool {
	fileName := getFileFromFilePath(path)
	fileSplit := strings.Split(fileName, ".")
	if len(fileSplit) == 2 && fileSplit[1] == "go" {
		return true
	}
	return false
}

func getFileFromFilePath(path string) string {
	splits := strings.Split(path, "/")
	index := len(splits) - 1
	return splits[index]
}

func getDirFromFilePath(path string) string {
	splits := strings.Split(path, "/")
	index := len(splits) - 2
	return splits[index]
}

func GetDirFromDirPath(path string) string {
	splits := strings.Split(path, "/")
	index := len(splits) - 1
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
