package teachercourse

import (
	"testing"

	"github.com/andamio-platform/transaction-specs/blob/main/classifier/utils"
)

func TestManageModules(t *testing.T) {
	accessTokenPolicy := "39b2876b2458b8cd869eb665b24740df6890684a3e6cd7ff6c28b84b"

	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	moduleScriptsV2PolicyId := "0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7"

	_, ok := ManageModules(tx, moduleScriptsV2PolicyId, accessTokenPolicy)
	t.Logf("ManageModules result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as ManageModules transaction")
	}
}
