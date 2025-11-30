package studentcourse

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestEnroll(t *testing.T) {

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds(courseStatePolicyIds)

	hashHex := "92d4ba99124b11aef75e4b2dd36e91b6e5b81e383c496836b6bdf3d9daf8dad2"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	_, ok := Enroll(tx)
	t.Logf("Enroll result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as Enroll transaction")
	}
}
