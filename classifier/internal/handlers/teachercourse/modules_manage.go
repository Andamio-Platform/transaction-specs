package teachercourse

import (
	"encoding/hex"
	"fmt"

	"github.com/Salvionied/apollo/serialization/PlutusData"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func ManageModules(tx *cardano.Tx, moduleScriptsV2PolicyId string) bool {

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
			return true
		}
	}

	return false
}
