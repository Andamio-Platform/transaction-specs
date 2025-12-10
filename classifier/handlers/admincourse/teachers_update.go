package admincourse

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"

	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/models"
)

func UpdateTeachers(tx *cardano.Tx) (*models.AdminCourseTeachersUpdate, bool) {
	var CourseGovernanceInput *cardano.TxOutput
	var CourseGovernanceOutput *cardano.TxOutput
	var courseID string
	cfg := config.Get()

	courseGovernanceV2Policy := cfg.CurrentV2().CourseGovernanceV2.MSCPolicyID
	outputs := tx.GetOutputs()
	for _, output := range outputs {
		multiAssets := output.GetAssets()
		for _, assets := range multiAssets {

			assetPolicyId := hex.EncodeToString(assets.GetPolicyId())
			if assetPolicyId == courseGovernanceV2Policy {
				courseID = hex.EncodeToString(assets.GetAssets()[0].GetName())
				CourseGovernanceOutput = output
			}

		}
	}

	inputs := tx.GetInputs()
	for _, input := range inputs {
		multiAssets := input.GetAsOutput().GetAssets()
		for _, assets := range multiAssets {
			assetPolicyId := hex.EncodeToString(assets.GetPolicyId())
			if assetPolicyId == courseGovernanceV2Policy {
				CourseGovernanceInput = input.GetAsOutput()
			}
		}
	}
	if CourseGovernanceInput == nil || CourseGovernanceOutput == nil {
		return nil, false
	} else {
		inputDatum := CourseGovernanceInput.GetDatum().GetPayload().GetArray().GetItems()
		outputDatum := CourseGovernanceOutput.GetDatum().GetPayload().GetArray().GetItems()

		oldTeachers := []string{}
		newTeachers := []string{}

		for _, inputItem := range inputDatum {
			oldTeachers = append(oldTeachers, string(inputItem.GetBoundedBytes()))
		}

		for _, outputItem := range outputDatum {
			newTeachers = append(newTeachers, string(outputItem.GetBoundedBytes()))
		}

		add := difference(newTeachers, oldTeachers)
		remove := difference(oldTeachers, newTeachers)

		return &models.AdminCourseTeachersUpdate{
			TxHash:   hex.EncodeToString(tx.GetHash()),
			CourseID: courseID,
			Add:      add,
			Remove:   remove,
		}, true
	}

}

func difference(a, b []string) []string {
	setB := make(map[string]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}

	diff := []string{}
	for _, v := range a {
		if _, found := setB[v]; !found {
			diff = append(diff, v)
		}
	}
	return diff
}
