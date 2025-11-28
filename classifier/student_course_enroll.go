package main

import (
	"encoding/hex"
	"slices"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func StudentCourseEnroll(tx *cardano.Tx, courseIds []string) bool {
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin == 1 {
					if slices.Contains(courseIds, hex.EncodeToString(mint.GetPolicyId())) {
						return true
					}
				}
			}
		}

	}
	return false
}
