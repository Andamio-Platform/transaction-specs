package studentcourse

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestClaimCredential(t *testing.T) {

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds(courseStatePolicyIds)

	hashHex := "75dcebcd0e768f93d6156dd174e072e1f281fa97ff3a47790fc7f198898dfdea"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	_, ok := ClaimCredential(config.Get(), tx)
	t.Logf("ClaimCredential result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as ClaimCredential transaction")
	}
}
