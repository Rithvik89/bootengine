package template

import (
	"testing"
)

func TestCreateDeepCopy(t *testing.T) {
	createDeepCopy("../../rithvik1", "../../rithvik2", true)
	// fmt.Println(entries)
}

func TestCloneProject(t *testing.T) {
	cloneProject("boot", "../../", "https://github.com/wangyoucao577/go-project-layout.git")
}

func TestRemoveNonEmptyDir(t *testing.T) {
	removeDir("../../boot")
	removeDir("../../cached")
}
