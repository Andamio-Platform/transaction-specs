package main

import (
	"encoding/json"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func TeacherCourseModulesManage(tx *cardano.Tx) bool {

	// Resolve Hash of Reference Inputs Scripts
	// Check if `0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7` is the hash

	referenceInputs := tx.GetReferenceInputs()
	referenceInputsJSON, _ := json.MarshalIndent(referenceInputs, "", "  ")
	println(string(referenceInputsJSON))

	return false

}
