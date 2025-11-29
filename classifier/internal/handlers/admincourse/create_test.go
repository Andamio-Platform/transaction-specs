package admincourse

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/classifier/internal/utils"
)

func TestCreateCourse(t *testing.T) {

	hashHex := "64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	LocalStateTokenPolicyId := "1b4d9c2a523f5042f3b188cedfe07aadee1151e418bf578819dc4b5a"
	InstanceGovernanceTokenPolicyId := "60e72e5ee056545fcb37f2d3f9b853daede356516ab5c80f886a652a"

	result := CreateCourse(tx, LocalStateTokenPolicyId, InstanceGovernanceTokenPolicyId)
	t.Logf("CreateCourse result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as CreateCourse transaction")
	}
}
