package template

import (
	"testing"
)

func TestCreateDeepCopy(t *testing.T) {
	createDeepCopy("../../rithvik1", "../../rithvik2", true)
	// fmt.Println(entries)
}

func TestCloneProject(t *testing.T) {
	cloneProject("boot", "../../")
}

func TestRemoveNonEmptyDir(t *testing.T) {
	removeDir("../../boot")
	removeDir("../../cached")
}
