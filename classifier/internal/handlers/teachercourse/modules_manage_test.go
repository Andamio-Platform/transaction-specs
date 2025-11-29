package teachercourse

import (
	"classifier/internal/utils"
	"testing"
)

func TestManageModules(t *testing.T) {

	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	moduleScriptsV2PolicyId := "0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7"

	result := ManageModules(tx, moduleScriptsV2PolicyId)
	t.Logf("ManageModules result: %v", result)

	if result != true {
		t.Error(hashHex + " should be classified as ManageModules transaction")
	}
}
