package studentcourse

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/classifier/internal/utils"
)

func TestClaimCredential(t *testing.T) {

	hashHex := "baf3d65fa644ce636536b0f9eef6591f26d2ee1561c26b80354cb17fb36a8eea"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}

	_, ok := ClaimCredential(tx, courseStatePolicyIds)
	t.Logf("ClaimCredential result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as ClaimCredential transaction")
	}
}
