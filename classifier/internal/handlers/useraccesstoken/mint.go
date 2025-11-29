package useraccesstoken

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func Mint(tx *cardano.Tx, accessTokenPolicy string) bool {
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin > 0 {
					if hex.EncodeToString(mint.GetPolicyId()) == accessTokenPolicy {
						return true
					}
				}
			}
		}

	}
	return false
}
