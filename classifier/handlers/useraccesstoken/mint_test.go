package useraccesstoken

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestMint(t *testing.T) {
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds([]string{})

	hashHex := "5b0bb8b17580e67a23c22d692d5f078daed7a19250684c9760b5d3bd64f70c3a"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	_, ok := Mint(tx)
	t.Logf("Mint result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as Mint transaction")
	}
}
