package admincourse

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"

	"github.com/andamio-platform/transaction-specs/classifier/internal/models"
)

func CreateCourse(tx *cardano.Tx, localStateTokenPolicy string, instanceGovernanceTokenPolicy string) (*models.AdminCourseCreate, bool) {
	isInitCourse := false

	requiredAssets := map[string]bool{
		// "LocalStateNFT":           false,	// TODO
		"LocalStateToken":                    false,
		"InstanceGovernanceToken":            false,
		"CourseStateScriptsV2ReferenceInput": false,
	}

	mints := tx.GetMint()

	// Check if CourseStateScriptsV2 is present
	referenceInputs := tx.GetReferenceInputs()
	for _, refInput := range referenceInputs {
		for _, assets := range refInput.GetAsOutput().GetAssets() {
			for _, asset := range assets.GetAssets() {
				assetName := string(asset.GetName())
				if assetName == "CourseStateScriptsV2" {
					requiredAssets["CourseStateScriptsV2ReferenceInput"] = true
				}
			}
		}
	}

	if len(mints) > 0 {
		for _, mint := range mints {
			if hex.EncodeToString(mint.GetPolicyId()) == localStateTokenPolicy {
				assets := mint.GetAssets()
				for _, asset := range assets {
					assetName := string(asset.GetName())
					if _, exists := requiredAssets[assetName]; exists {
						if asset.GetMintCoin() > 0 {
							requiredAssets[assetName] = true
						}
					}
				}
			}

			if hex.EncodeToString(mint.GetPolicyId()) == instanceGovernanceTokenPolicy {
				requiredAssets["InstanceGovernanceToken"] = true
			}
		}

		// Check if all required assets are present
		allFound := true
		for _, found := range requiredAssets {
			if !found {
				allFound = false
				break
			}
		}
		isInitCourse = allFound
	}

	if isInitCourse {
		return &models.AdminCourseCreate{
			TxHash: hex.EncodeToString(tx.GetHash()),
			// TODO: Extract other fields
		}, true
	}

	return nil, false
}
