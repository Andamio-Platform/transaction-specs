package main

import (
	"testing"
)

func TestAdminCourseCreate(t *testing.T) {

	hashHex := "64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc"
	tx := GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	result := AdminCourseCreate(tx)
	t.Logf("AdminCourseCreate result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as AdminCourseCreate transaction")
	}
}
