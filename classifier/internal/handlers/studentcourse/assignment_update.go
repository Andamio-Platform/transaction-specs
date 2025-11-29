package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func UpdateAssignment(tx *cardano.Tx, courseStatePolicyIds []string) bool {
	var oldContent string
	var updatedContent string

	// TODO: Check if courseStateToken is in inputs and has datum with constructor 1

	outputs := tx.GetOutputs()
	for _, output := range outputs {
		multiassets := output.GetAssets()
		for _, ma := range multiassets {
			if slices.Contains(courseStatePolicyIds, hex.EncodeToString(ma.GetPolicyId())) {
				datum := output.GetDatum().GetPayload()
				// Check if constructor 1
				if datum.GetConstr().GetTag() == 122 {
					content := datum.GetConstr().GetFields()[1].GetBoundedBytes()
					updatedContent = hex.EncodeToString(content)
				}
			}
		}
	}

	return oldContent != updatedContent
}
