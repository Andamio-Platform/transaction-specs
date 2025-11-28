package main

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func AdminCourseCreate(tx *cardano.Tx) bool {
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
			if hex.EncodeToString(mint.GetPolicyId()) == "1b4d9c2a523f5042f3b188cedfe07aadee1151e418bf578819dc4b5a" {
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

			if hex.EncodeToString(mint.GetPolicyId()) == "60e72e5ee056545fcb37f2d3f9b853daede356516ab5c80f886a652a" {
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
