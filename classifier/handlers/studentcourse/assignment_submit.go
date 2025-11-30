package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func SubmitAssignment(tx *cardano.Tx) (*models.StudentCourseAssignmentSubmit, bool) {
	courseStatePolicyIds := config.GetCourseStatePolicyIds()

	// TODO: Check if courseStateToken is minted (OR) input courseStateToken has datum with constructor 0

	outputs := tx.GetOutputs()
	for _, output := range outputs {
		multiassets := output.GetAssets()
		for _, ma := range multiassets {
			if slices.Contains(courseStatePolicyIds, hex.EncodeToString(ma.GetPolicyId())) {
				datum := output.GetDatum().GetPayload()
				// Check if constructor 1
				if datum.GetConstr().GetTag() == 122 {
					// Check if content is not empty string
					content := datum.GetConstr().GetFields()[1].GetBoundedBytes()
					if len(content) > 0 {

						alias := string(ma.GetAssets()[0].GetName())

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

						assignmentID := hex.EncodeToString(datum.GetConstr().GetFields()[0].GetBoundedBytes())

						return &models.StudentCourseAssignmentSubmit{
							TxHash:       hex.EncodeToString(tx.GetHash()),
							Alias:        alias,
							CourseID:     courseID,
							AssignmentID: assignmentID,
							Content:      hex.EncodeToString(content),
						}, true
					}
				}
			}
		}
	}

	return nil, false
}
