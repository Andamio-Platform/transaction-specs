package useraccesstoken

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/blob/main/classifier/utils"
)

func TestMint(t *testing.T) {

	hashHex := "5b0bb8b17580e67a23c22d692d5f078daed7a19250684c9760b5d3bd64f70c3a"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	accessTokenPolicy := "39b2876b2458b8cd869eb665b24740df6890684a3e6cd7ff6c28b84b"

	_, ok := Mint(tx, accessTokenPolicy)
	t.Logf("Mint result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as Mint transaction")
	}
}
