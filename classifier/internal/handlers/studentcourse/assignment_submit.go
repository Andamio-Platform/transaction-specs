package studentcourse

import (
	"encoding/hex"
	"slices"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func SubmitAssignment(tx *cardano.Tx, courseStatePolicyIds []string) bool {

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
						return true
					}
				}
			}
		}
	}

	return false
}
