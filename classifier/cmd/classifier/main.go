package main

import (
	"classifier/internal/handlers/teachercourse"
	"classifier/internal/utils"
)

func main() {
	hashHex := "7a96d45238788c92143a3bc2aaae2d405c25efe5da1281b2930bd42e717d90fa"
	tx := utils.GetCardanoTx(hashHex)

	accessTokenPolicy := "39b2876b2458b8cd869eb665b24740df6890684a3e6cd7ff6c28b84b"
	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}

	teachercourse.AssessAssignments(tx, accessTokenPolicy, courseStatePolicyIds)
}
