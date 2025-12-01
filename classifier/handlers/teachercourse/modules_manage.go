package teachercourse

import (
	"encoding/hex"
	"fmt"

	"github.com/Salvionied/apollo/serialization/PlutusData"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func ManageModules(tx *cardano.Tx) (*models.TeacherCourseModulesManage, bool) {
	moduleScriptsV2PolicyId := config.Get().CurrentV2().ModuleScriptsV2.MSCPolicyID
	accessTokenPolicy := config.Get().CurrentV2().IndexMS.MSCPolicyID

	referenceInputs := tx.GetReferenceInputs()
	for _, refInput := range referenceInputs {
		out := refInput.GetAsOutput()
		if out == nil {
			continue
		}

		script := out.GetScript()
		if script == nil || script.GetPlutusV3() == nil {
			continue
		}

		plutusV3 := PlutusData.PlutusV3Script(script.GetPlutusV3())

		hash, err := plutusV3.Hash()
		if err != nil {
			fmt.Println("Error computing hash:", err)
			continue
		}

		if hex.EncodeToString(hash.Bytes()) == moduleScriptsV2PolicyId {
			var alias string
			var courseID string
			outputs := tx.GetOutputs()
			for _, output := range outputs {
				multiassets := output.GetAssets()
				for _, multiasset := range multiassets {
					assets := multiasset.GetAssets()
					for _, asset := range assets {
						if hex.EncodeToString(multiasset.GetPolicyId()) == accessTokenPolicy {
							alias = string(asset.GetName())[1:]
						}
						if hex.EncodeToString(multiasset.GetPolicyId()) == moduleScriptsV2PolicyId {
							datum := output.GetDatum().GetPayload()
							courseID = hex.EncodeToString(datum.GetConstr().GetFields()[0].GetBoundedBytes())
						}
					}
				}
			}

			var modulesCreated []models.ModulesCreated
			mints := tx.GetMint()
			for _, mint := range mints {
				if hex.EncodeToString(mint.GetPolicyId()) == moduleScriptsV2PolicyId {
					assets := mint.GetAssets()
					for _, asset := range assets {
						var slts models.StringArray
						var prerequisites models.StringArray = models.StringArray{}
						if asset.GetMintCoin() > 0 {
							assignmentID := hex.EncodeToString(asset.GetName())
							redeemer := mint.GetRedeemer().GetPayload()
							modules := redeemer.GetConstr().GetFields()[2].GetArray().GetItems()
							sltsPlutusData := modules[0].GetArray().GetItems()
							for _, sltPlutusData := range sltsPlutusData {
								slts = append(slts, string(sltPlutusData.GetBoundedBytes()))
							}
							if len(modules) > 1 {
								prerequisitesPlutusData := modules[1].GetArray().GetItems()
								for _, prerequisitePlutusData := range prerequisitesPlutusData {
									prerequisites = append(prerequisites, string(prerequisitePlutusData.GetBoundedBytes()))
								}
							}
							modulesCreated = append(modulesCreated, models.ModulesCreated{
								AssignmentID: assignmentID,
								Module: models.ModuleCreate{
									SLTs:          slts,
									Prerequisites: prerequisites,
								},
							})
						}
					}
				}
			}

			return &models.TeacherCourseModulesManage{
				TxHash:   hex.EncodeToString(tx.GetHash()),
				Alias:    alias,
				CourseID: courseID,
				Modules: models.Modules{
					Create: modulesCreated,
					Update: []models.ModulesUpdated{}, // TODO: Implement
					Delete: []models.ModulesDeleted{}, // TODO: Implement
				},
			}, true
		}
	}

	return nil, false
}
