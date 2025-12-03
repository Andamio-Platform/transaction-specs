package admincourse

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestCreateCourse(t *testing.T) {

	// instanceStakingCredential := "99247f1925b72a641a0e3ade6191bdbe7b3e3fc1fd45f4ae0f4a54b6"

	hashHex := "64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	// Initialize config
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds([]string{})

	_, ok := CreateCourse(tx)
	t.Logf("CreateCourse result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as CreateCourse transaction")
	}
}
