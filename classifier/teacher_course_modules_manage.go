package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Salvionied/apollo/serialization/PlutusData"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func TeacherCourseModulesManage(tx *cardano.Tx) bool {
	targetHash := "0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7"

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

		if hex.EncodeToString(hash.Bytes()) == targetHash {
			return true
		}
	}

	return false
}
