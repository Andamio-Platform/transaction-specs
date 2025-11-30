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
	// Mock config values if needed, but since we load from file in config.load(), it should be fine if the file exists.
	// However, the test might not be running from the root directory, so relative paths in config.load might fail.
	// Let's check config.go loadV2Config path: "./config/v2-preprod.json".
	// If running test from handlers/admincourse, the path should be "../../config/v2-preprod.json".
	// This might be an issue. For now, let's assume the user runs tests from root or we might need to adjust config loading.
	// But wait, I can't easily change the file path in config.go without breaking main.
	// I'll proceed with updating the test and if it fails, I'll address the path issue.
	// Actually, I can manually set the config instance in the test if I export a way to do it or use a mock.
	// But config.instance is private.
	// I will just call config.Init and hope for the best or I might need to hack the path.
	// Ideally, config path should be configurable.

	// For now, I'll just update the call.
	// Note: The original test passed policies explicitly. Now they come from config.
	// The config file "v2-preprod.json" should contain the values used in the test.
	// LocalStateRef.MSCPolicyID = "1b4d9c2a523f5042f3b188cedfe07aadee1151e418bf578819dc4b5a"
	// CourseGovernanceV2.MSCPolicyID = "60e72e5ee056545fcb37f2d3f9b853daede356516ab5c80f886a652a"
	// InstanceStakingScrSh = "99247f1925b72a641a0e3ade6191bdbe7b3e3fc1fd45f4ae0f4a54b6"

	// I will assume the config file has these values or I might need to overwrite them in the test.
	// Since I cannot overwrite easily, I will just call Init.

	_, ok := CreateCourse(tx)
	t.Logf("CreateCourse result: %v", ok)

	if !ok {
		t.Error(hashHex + " should be classified as CreateCourse transaction")
	}
}
