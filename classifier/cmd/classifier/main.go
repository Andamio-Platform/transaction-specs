package main

import (
	"classifier/internal/handlers/admincourse"
	"classifier/internal/handlers/studentcourse"
	"classifier/internal/handlers/teachercourse"
	"classifier/internal/utils"
	"fmt"
)

func main() {
	txHashes := []string{
		"35d93ccfe17ccd6de427c66818f19eb79d729b7abd825be02441a70dfd769aff",
		"5b0bb8b17580e67a23c22d692d5f078daed7a19250684c9760b5d3bd64f70c3a",
		"fbc7d62489b51e81026b2ed695417cba6a657ac5429d4bb8211bfc7d5aa667a9",
		"c58ddc70c8f322937d14a41e140bcbd342e9be1e96f0ba08005d1dcdb9540654",
		"289c60362495eed68e209a59ff002972cf22969207854059afce2f8ac8576354",
		"ff695f713489b7b7588f18b66f7a2744455a5bd474845153331ae845dbe425a0",
		"b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d",
		"92d4ba99124b11aef75e4b2dd36e91b6e5b81e383c496836b6bdf3d9daf8dad2",
		"7a96d45238788c92143a3bc2aaae2d405c25efe5da1281b2930bd42e717d90fa",
		"baf3d65fa644ce636536b0f9eef6591f26d2ee1561c26b80354cb17fb36a8eea",
		"81a7be23da69be08b330b021a3b7d435e2a726cfceb51bb0c93ffa2c3763ce45",
		"80fac6d6516429700728e1a4883eac3f19eea46fc1ebf2810dd01359c9e346fe",
		"863e38af684c603900a3297ed726ddd1f8def8dce68180a73e2267a4f69cf104",
		"1b5063ba14cacd6debaf6f4654a44711b22ea6d46482d94d8fa5e01dc02edb36",
		"75dcebcd0e768f93d6156dd174e072e1f281fa97ff3a47790fc7f198898dfdea",
		"64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc",
	}

	// Policies for the handlers
	localStateTokenPolicy := "39b2876b2458b8cd869eb665b24740df6890684a3e6cd7ff6c28b84b"
	instanceGovernanceTokenPolicy := "d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0"
	courseStatePolicyIds := []string{
		"d8475bbfe87cdd18592b8d0c623be1d9be961ed93f75ded26b00e9b0",
	}

	for _, hash := range txHashes {
		fmt.Println("Processing transaction:", hash)

		tx := utils.GetCardanoTx(hash)
		if tx == nil {
			fmt.Println("Failed to fetch transaction")
			continue
		}

		classified := false

		// ===== AdminCourse =====
		if admincourse.CreateCourse(tx, localStateTokenPolicy, instanceGovernanceTokenPolicy) {
			fmt.Println("Handler: AdminCourse → CREATE_COURSE")
			classified = true
		}

		// ===== TeacherCourse =====
		if teachercourse.AssessAssignments(tx, localStateTokenPolicy, courseStatePolicyIds) {
			fmt.Println("Handler: TeacherCourse → ASSESS_ASSIGNMENTS")
			classified = true
		}

		// ===== StudentCourse =====
		if studentcourse.Enroll(tx, courseStatePolicyIds) {
			fmt.Println("Handler: StudentCourse → ENROLL")
			classified = true
		}

		if !classified {
			fmt.Println("No handler claimed this transaction → REFUSE / INVALID")
		}

		fmt.Println("-------------------------------------------------")
	}
}
