package teachercourse

import (
	"testing"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func TestAssessAssignments(t *testing.T) {

	hashHex := "7a96d45238788c92143a3bc2aaae2d405c25efe5da1281b2930bd42e717d90fa"
	tx := utils.GetCardanoTx(hashHex)

	if tx == nil {
		t.Fatal("Failed to retrieve transaction")
	}

	courseStatePolicyIds := []string{"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"}
	config.Init(constants.PREPROD)
	config.SetCourseStatePolicyIds(courseStatePolicyIds)

	_, ok := AssessAssignments(config.Get(), tx)
	t.Logf("AssessAssignments result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as AssessAssignments transaction")
	}
}
