package main

import (
	"testing"
)

func TestTeacherCourseModulesManage(t *testing.T) {

	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	result := TeacherCourseModulesManage(tx)
	t.Logf("TeacherCourseModulesManage result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as TeacherCourseModulesManage transaction")
	}
}
