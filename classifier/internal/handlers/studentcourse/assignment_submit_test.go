package studentcourse

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/classifier/internal/utils"
)

func TestSubmitAssignment(t *testing.T) {

	hashHex := "863e38af684c603900a3297ed726ddd1f8def8dce68180a73e2267a4f69cf104"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}

	result := SubmitAssignment(tx, courseStatePolicyIds)
	t.Logf("SubmitAssignment result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as SubmitAssignment transaction")
	}
}
