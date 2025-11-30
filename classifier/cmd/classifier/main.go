package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/handlers/admincourse"
	"github.com/andamio-platform/transaction-specs/classifier/handlers/studentcourse"
	"github.com/andamio-platform/transaction-specs/classifier/handlers/teachercourse"
	"github.com/andamio-platform/transaction-specs/classifier/handlers/useraccesstoken"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func main() {
	txHashes := []string{
		// "35d93ccfe17ccd6de427c66818f19eb79d729b7abd825be02441a70dfd769aff", // protocol init
		// "5b0bb8b17580e67a23c22d692d5f078daed7a19250684c9760b5d3bd64f70c3a", // UserAccessToken → MINT
		// "fbc7d62489b51e81026b2ed695417cba6a657ac5429d4bb8211bfc7d5aa667a9", // protocol init
		// "c58ddc70c8f322937d14a41e140bcbd342e9be1e96f0ba08005d1dcdb9540654", // UserAccessToken → MINT
		// "289c60362495eed68e209a59ff002972cf22969207854059afce2f8ac8576354", // UserAccessToken → MINT
		// "ff695f713489b7b7588f18b66f7a2744455a5bd474845153331ae845dbe425a0", // UserAccessToken → MINT
		"b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d", // TeacherCourse → MANAGE_MODULES
		// "92d4ba99124b11aef75e4b2dd36e91b6e5b81e383c496836b6bdf3d9daf8dad2", // StudentCourse → ENROLL , StudentCourse → SUBMIT_ASSIGNMENT
		// "7a96d45238788c92143a3bc2aaae2d405c25efe5da1281b2930bd42e717d90fa", // TeacherCourse → ASSESS_ASSIGNMENTS
		// "baf3d65fa644ce636536b0f9eef6591f26d2ee1561c26b80354cb17fb36a8eea", // StudentCourse → CLAIM_CREDENTIAL
		// "81a7be23da69be08b330b021a3b7d435e2a726cfceb51bb0c93ffa2c3763ce45", // TeacherCourse → MANAGE_MODULES
		// "80fac6d6516429700728e1a4883eac3f19eea46fc1ebf2810dd01359c9e346fe", // UserAccessToken → MINT
		// "863e38af684c603900a3297ed726ddd1f8def8dce68180a73e2267a4f69cf104", // StudentCourse → ENROLL , StudentCourse → SUBMIT_ASSIGNMENT
		// "1b5063ba14cacd6debaf6f4654a44711b22ea6d46482d94d8fa5e01dc02edb36", // TeacherCourse → ASSESS_ASSIGNMENTS
		// "75dcebcd0e768f93d6156dd174e072e1f281fa97ff3a47790fc7f198898dfdea", // StudentCourse → CLAIM_CREDENTIAL
		// "64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc", // AdminCourse → CREATE_COURSE
	}

	network := constants.PREPROD
	instanceStakingCredential := "99247f1925b72a641a0e3ade6191bdbe7b3e3fc1fd45f4ae0f4a54b6"

	// Policies
	accessTokenPolicy := "39b2876b2458b8cd869eb665b24740df6890684a3e6cd7ff6c28b84b"
	localStateReferencePolicy := "1b4d9c2a523f5042f3b188cedfe07aadee1151e418bf578819dc4b5a"
	courseGovernanceV2Policy := "60e72e5ee056545fcb37f2d3f9b853daede356516ab5c80f886a652a"
	courseStatePolicyIds := []string{
		"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0",
	}
	moduleScriptsV2PolicyId := "0881d005d4301748df5aab08fbd302ad62f06a1b6b154664c96b9ba7"

	for _, hash := range txHashes {
		fmt.Println("Processing transaction:", hash)

		tx := utils.GetCardanoTx(hash)
		if tx == nil {
			fmt.Println("Failed to fetch transaction")
			continue
		}

		var matched []string

		// ===== UserAccessToken =====
		if model, ok := useraccesstoken.Mint(tx, accessTokenPolicy); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("UserAccessToken → MINT\n%s", string(jsonBytes)))
		}

		// ===== AdminCourse =====
		if model, ok := admincourse.CreateCourse(tx, localStateReferencePolicy, courseGovernanceV2Policy, instanceStakingCredential, network); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("AdminCourse → CREATE_COURSE\n%s", string(jsonBytes)))
		}
		if model, ok := admincourse.UpdateTeachers(tx); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("AdminCourse → UPDATE_TEACHERS\n%s", string(jsonBytes)))
		}

		// ===== TeacherCourse =====
		if model, ok := teachercourse.AssessAssignments(tx, accessTokenPolicy, courseStatePolicyIds); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("TeacherCourse → ASSESS_ASSIGNMENTS\n%s", string(jsonBytes)))
		}
		if model, ok := teachercourse.ManageModules(tx, moduleScriptsV2PolicyId, accessTokenPolicy); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("TeacherCourse → MANAGE_MODULES\n%s", string(jsonBytes)))
		}

		// ===== StudentCourse =====
		if model, ok := studentcourse.Enroll(tx, courseStatePolicyIds); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("StudentCourse → ENROLL\n%s", string(jsonBytes)))
		}
		if model, ok := studentcourse.SubmitAssignment(tx, courseStatePolicyIds); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("StudentCourse → SUBMIT_ASSIGNMENT\n%s", string(jsonBytes)))
		}
		if model, ok := studentcourse.UpdateAssignment(tx, courseStatePolicyIds); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("StudentCourse → UPDATE_ASSIGNMENT\n%s", string(jsonBytes)))
		}
		if model, ok := studentcourse.ClaimCredential(tx, courseStatePolicyIds); ok {
			jsonBytes, _ := json.MarshalIndent(model, "", "  ")
			matched = append(matched, fmt.Sprintf("StudentCourse → CLAIM_CREDENTIAL\n%s", string(jsonBytes)))
		}

		if len(matched) == 0 {
			fmt.Println("No handler claimed this transaction → REFUSE / INVALID")
		} else if len(matched) == 1 {
			fmt.Println("Classified as:", matched[0])
		} else {
			fmt.Printf("⚠️  Multiple classifications (%d):\n%s\n", len(matched), strings.Join(matched, "\n-------------------\n"))
		}

		fmt.Println("-------------------------------------------------")
	}
}
