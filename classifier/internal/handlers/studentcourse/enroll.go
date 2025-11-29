package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/andamio-platform/transaction-specs/classifier/internal/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func Enroll(tx *cardano.Tx, courseStatePolicyIds []string) (*models.StudentCourseEnroll, bool) {
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin == 1 {
					if slices.Contains(courseStatePolicyIds, hex.EncodeToString(mint.GetPolicyId())) {
						return &models.StudentCourseEnroll{
							TxHash: hex.EncodeToString(tx.GetHash()),
							// TODO: Extract other fields
						}, true
					}
				}
			}
		}

	}
	return nil, false
}
