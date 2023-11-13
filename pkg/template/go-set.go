package template

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func isFeasibleToCreateProject(path string, name string) bool {
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("path invalid: %s", err)
		return false
	} else {
		// check what if the path mentioned is of a file return false.
		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if !fileInfo.IsDir() {
			fmt.Print("cannot create a project here as mentioned path is of a file")
			return false
		}
		files, err := file.Readdir(-1)
		for _, f := range files {
			if name == f.Name() {
				fmt.Printf("Dir with name %s already exits", f.Name())
				return false
			}
		}
		return true
	}
}

func CloneProject(projectName string, DestinationPath string, url string) {

	// 1. Pick the name of the project from options and AbsPath.
	// 2. Create a directory in the name mentioned above and initialize a git repo.

	// 3. Load the respective language repo from github.
	// currentpath, err := os.Getwd()
	// projectName is where we clone our project
	path := DestinationPath + "/" + projectName

	// TODO: check for scenerios like null strings in Projectname,destinationpath,url
	if url == "" {
		url = "https://github.com/wangyoucao577/go-project-layout.git"
	}
	if DestinationPath == "" || projectName == "" {
		panic("Use valid Destination path and project name")
	}

	// *check if there is some project with the projectname as mentioned above in the dest path.

	if !isFeasibleToCreateProject(DestinationPath, projectName) {
		return
	}

	// TODO: instead of clone in onto disk why can't we try in memory clone ?.
	cachedPath := "./cached"

	_, err := git.PlainClone(cachedPath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Error initializing repository: %v\n", err)
		return
	}

	if err != nil {
		panic("Not able to clone content into cached repo.")
	}

	createDeepCopy(cachedPath, path, true)

	// 4. Prepare it as per language run book.
	// 5. Move the contents into user-mentioned folder.

	// 6. remove the cached folder
	os.RemoveAll(cachedPath)

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

func createFile(path string) {
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

func removeDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Successfully deleted at %s", path)
}

// SEPERATOR SHOULD BE DIFFERENT FOR WINDOWS AND LINUX BASED SYSTEMS.

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

func isGoFile(path string) bool {
	fileName := getFileFromPath(path)
	fileSplit := strings.Split(fileName, ".")
	if len(fileSplit) == 2 && fileSplit[1] == "go" {
		return true
	}
	return false
}
