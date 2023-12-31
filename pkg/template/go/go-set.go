package go_template

import (
	"bootengine/utils"
	"fmt"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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

	repo := createDeepCopy(cachedPath, path, true)
	// git add and git commit the repo.
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}
	commit, err := worktree.Commit("Project Layout Structured by devboot", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "devboot",
			Email: "root@devboot.in",
			When:  time.Now(),
		},
	})

	fmt.Println("Changes committed. Commit Hash:", commit)

	// 4. Prepare it as per language run book.
	// 5. Move the contents into user-mentioned folder.

	// 6. remove the cached folder
	os.RemoveAll(cachedPath)

}

func createDeepCopy(source string, destination string, root bool) *git.Repository {
	files, err := os.Open(source) //open the directory to read files in the directory
	if err != nil {
		fmt.Println("error opening directory:", err) //print error if directory is not opened
		return nil
	}

	defer files.Close() //close the directory opened

	fileInfo, err := files.Stat()
	if !fileInfo.IsDir() {
		utils.CreateFile(destination)
		return nil
	}

	fileInfos, err := files.Readdir(-1) //read the files from the directory
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return nil
	}

	if utils.GetDirFromDirPath(destination) == ".git" {
		return nil
	}

	repo := utils.CreateFolder(destination, root)

	for _, fileInfos := range fileInfos {
		subPathDestination := destination + "/" + fileInfos.Name()
		subPathSource := source + "/" + fileInfos.Name()
		createDeepCopy(subPathSource, subPathDestination, false)
	}
	return repo
}
