package admincourse

import (
	"encoding/hex"
	"fmt"

	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/plutusData"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"

	"github.com/andamio-platform/transaction-specs/classifier/models"
)

func CreateCourse(tx *cardano.Tx) (*models.AdminCourseCreate, bool) {
	isInitCourse := false

	localStateReferencePolicy := config.Get().CurrentV2().LocalStateRef.MSCPolicyID
	courseGovernanceV2Policy := config.Get().CurrentV2().CourseGovernanceV2.MSCPolicyID
	instanceStakingScriptHash := config.Get().CurrentV2().InstanceStakingScrSh
	network := config.Get().Network

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
			if hex.EncodeToString(mint.GetPolicyId()) == localStateReferencePolicy {
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

			if hex.EncodeToString(mint.GetPolicyId()) == courseGovernanceV2Policy {
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
		var courseStateScript *cardano.Script
		var admin string
		var courseID string
		var teachers []string

		outputs := tx.GetOutputs()
		for _, output := range outputs {
			multiAssets := output.GetAssets()
			for _, assets := range multiAssets {
				for _, asset := range assets.GetAssets() {
					assetPolicyId := hex.EncodeToString(assets.GetPolicyId())
					assetName := string(asset.GetName())
					if assetPolicyId == localStateReferencePolicy && assetName == "LocalStateToken" {
						courseStateScript = output.GetScript()
						datum := output.GetDatum().GetPayload()
						courseID = hex.EncodeToString(datum.GetBoundedBytes())
					}
					if assetName == "LocalStateNFT" { // TODO: add policy id check
						datum := output.GetDatum().GetPayload()
						admin = string(datum.GetBoundedBytes())
					}
					if assetPolicyId == courseGovernanceV2Policy {
						datum := output.GetDatum().GetPayload()
						teachersPlutusData := datum.GetArray().GetItems()
						for _, teacherPlutusData := range teachersPlutusData {
							teacher := teacherPlutusData.GetBoundedBytes()
							teachers = append(teachers, string(teacher))
						}
					}
				}
			}
		}

		plutusV3 := plutusData.PlutusV3Script(courseStateScript.GetPlutusV3())

		courseStatePolicyId, err := plutusV3.Hash()
		if err != nil {
			fmt.Println("Error computing hash:", err)
		}
		stakingCredential, err := hex.DecodeString(instanceStakingScriptHash)
		if err != nil {
			fmt.Println("Error decoding staking credential:", err)
		}
		courseAddress := plutusV3.ToAddress(stakingCredential, true, network)

		return &models.AdminCourseCreate{
			TxHash:              hex.EncodeToString(tx.GetHash()),
			CourseID:            courseID,
			Admin:               admin,
			Teachers:            teachers,
			CourseAddress:       courseAddress.String(),
			CourseStatePolicyId: hex.EncodeToString(courseStatePolicyId.Bytes()),
		}, true
	}

	return nil, false
}
