package go_template

import (
	"bootengine/utils"
	"testing"
)

func TestCreateDeepCopy(t *testing.T) {
	createDeepCopy("../../rithvik1", "../../rithvik2", true)
}

func TestCloneProject(t *testing.T) {
	CloneProject("rithu", "/home/ra/Documents", "https://github.com/wangyoucao577/go-project-layout.git")
}

func TestRemoveNonEmptyDir(t *testing.T) {
	utils.RemoveDir("../../boot")
	utils.RemoveDir("../../cached")
}
