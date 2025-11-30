package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func ClaimCredential(tx *cardano.Tx, courseStatePolicyIds []string) (*models.StudentCourseCredentialClaim, bool) {
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin == -1 {
					if slices.Contains(courseStatePolicyIds, hex.EncodeToString(mint.GetPolicyId())) {

						redeemer := mint.GetRedeemer().GetPayload()

						alias := string(redeemer.GetConstr().GetFields()[0].GetBoundedBytes())

						var credentials []string
						credentialsPlutusData := redeemer.GetConstr().GetFields()[1].GetArray().GetItems()
						for _, credentialPlutusData := range credentialsPlutusData {
							credentials = append(credentials, hex.EncodeToString(credentialPlutusData.GetBoundedBytes()))
						}

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
						return &models.StudentCourseCredentialClaim{
							TxHash:      hex.EncodeToString(tx.GetHash()),
							Alias:       alias,
							CourseID:    courseID,
							Credentials: credentials,
						}, true
					}
				}
			}
		}

	}
	return nil, false
}
