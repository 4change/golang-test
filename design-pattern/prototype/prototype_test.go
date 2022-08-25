package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	resume := NewResume()

	resume.setPersonalInfo("hclA", "男", "22")
	resume.setWorkExperience("3", "Apple")
	resume.display()

	cloneResume := resume.clone()
	cloneResume.setPersonalInfo("hclB", "女", "22")
	cloneResume.setWorkExperience("3", "HW")
	cloneResume.display()
}
