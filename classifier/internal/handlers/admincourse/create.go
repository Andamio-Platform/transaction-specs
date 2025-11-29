package admincourse

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func CreateCourse(tx *cardano.Tx, LocalStateTokenPolicyId string, InstanceGovernanceTokenPolicyId string) bool {
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
			if hex.EncodeToString(mint.GetPolicyId()) == LocalStateTokenPolicyId {
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

			if hex.EncodeToString(mint.GetPolicyId()) == InstanceGovernanceTokenPolicyId {
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

	return isInitCourse
}
