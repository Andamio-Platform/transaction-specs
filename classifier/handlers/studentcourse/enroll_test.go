package studentcourse

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/blob/main/classifier/utils"
)

func TestEnroll(t *testing.T) {

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}

	hashHex := "92d4ba99124b11aef75e4b2dd36e91b6e5b81e383c496836b6bdf3d9daf8dad2"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	_, ok := Enroll(tx, courseStatePolicyIds)
	t.Logf("Enroll result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as Enroll transaction")
	}
}
