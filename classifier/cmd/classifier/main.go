package main

import (
	"classifier/internal/handlers/teachercourse"
	"classifier/internal/utils"
)

func main() {
	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := utils.GetCardanoTx(hashHex)

	moduleScriptsV2PolicyId := "0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7"
	teachercourse.ManageModules(tx, moduleScriptsV2PolicyId)
}
