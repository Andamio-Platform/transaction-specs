package teachercourse

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestManageModules(t *testing.T) {
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds([]string{})

	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	_, ok := ManageModules(config.Get(), tx)
	t.Logf("ManageModules result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as ManageModules transaction")
	}
}
