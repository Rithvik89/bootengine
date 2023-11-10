package template

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func cloneProject(projectName string, RelPath string) {

	// 1. Pick the name of the project from options and AbsPath.
	// 2. Create a directory in the name mentioned above and initialize a git repo.

	// 3. Load the respective language repo from github.
	// currentpath, err := os.Getwd()
	// projectName is where we clone our project
	path := RelPath + projectName
	cachedPath := "../../cached"

	_, err := git.PlainClone(cachedPath, false, &git.CloneOptions{
		URL:      "https://github.com/wangyoucao577/go-project-layout.git",
		Progress: os.Stdout,
	})

	if err != nil {
		panic("Not able to clone content")
	}

	createDeepCopy(cachedPath, path, true)

	// 4. Prepare it as per language run book.
	// 5. Move the contents into user-mentioned folder.

	// 6. remove the cached folder
	os.Remove(cachedPath)

}

func createDeepCopy(source string, destination string, root bool) {
	files, err := os.Open(source) //open the directory to read files in the directory
	if err != nil {
		fmt.Println("error opening directory:", err) //print error if directory is not opened
		return
	}

	defer files.Close() //close the directory opened

	fileInfo, err := files.Stat()
	if !fileInfo.IsDir() {
		createFile(destination)
		return
	}

	fileInfos, err := files.Readdir(-1) //read the files from the directory
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return
	}

	createFolder(destination, root)

	for _, fileInfos := range fileInfos {
		subPathDestination := destination + "/" + fileInfos.Name()
		subPathSource := source + "/" + fileInfos.Name()
		createDeepCopy(subPathSource, subPathDestination, false)
	}
	return
}

func createFolder(path string, root bool) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Printf("Not able to create a directory : %s", err)
		return
	}
	fmt.Printf("Directory created at %s", path)
}

func createFile(path string) {
	_, err := os.Create(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("File Created at %s", path)
}

func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Successfully deleted at %s", path)
}
