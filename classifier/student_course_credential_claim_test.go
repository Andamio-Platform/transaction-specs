package main

import (
	"testing"
)

func TestStudentCourseCredentialClaim(t *testing.T) {

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}

	hashHex := "baf3d65fa644ce636536b0f9eef6591f26d2ee1561c26b80354cb17fb36a8eea"
	tx := GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	result := StudentCourseCredentialClaim(tx, courseStatePolicyIds)
	t.Logf("StudentCourseCredentialClaim result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as StudentCourseCredentialClaim transaction")
	}
}
