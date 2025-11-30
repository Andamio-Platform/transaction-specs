package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func Enroll(tx *cardano.Tx, courseStatePolicyIds []string) (*models.StudentCourseEnroll, bool) {
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin == 1 {
					if slices.Contains(courseStatePolicyIds, hex.EncodeToString(mint.GetPolicyId())) {

						alias := string(asset.GetName())

						var courseID string
						referenceInputs := tx.GetReferenceInputs()
						for _, refInput := range referenceInputs {
							refOutput := refInput.GetAsOutput()
							refMultiassets := refOutput.GetAssets()
							for _, refMa := range refMultiassets {
								for _, refAsset := range refMa.GetAssets() {
									if string(refAsset.GetName()) == "LocalStateToken" {
										datum := refOutput.GetDatum().GetPayload()
										courseID = hex.EncodeToString(datum.GetBoundedBytes())
									}
								}
							}
						}

						return &models.StudentCourseEnroll{
							TxHash:   hex.EncodeToString(tx.GetHash()),
							Alias:    alias,
							CourseID: courseID,
						}, true
					}
				}
			}
		}

	}
	return nil, false
}
