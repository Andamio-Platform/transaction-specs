package teachercourse

import (
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func AssessAssignments(tx *cardano.Tx) (*models.TeacherCourseAssignmentsAssess, bool) {
	cfg := config.Get()
	accessTokenPolicy := cfg.CurrentV2().IndexMS.MSCPolicyID
	courseStatePolicyIds := config.GetCourseStatePolicyIds()

	// Check outputs has one of the courseStatePolicyIds
	// Check if user (u) token is different with course state token
	var userToken string
	type Assessment struct {
		CourseStateToken       string
		CourseStateTokenOutput *cardano.TxOutput
	}
	var assessments []Assessment

	outputs := tx.GetOutputs()
	for _, output := range outputs {
		multiassets := output.GetAssets()
		for _, multiasset := range multiassets {
			policyId := multiasset.GetPolicyId()
			if hex.EncodeToString(policyId) == accessTokenPolicy {
				userToken = string(multiasset.GetAssets()[0].GetName())
			}
			if slices.Contains(courseStatePolicyIds, hex.EncodeToString(policyId)) {
				asset := multiasset.GetAssets()
				for _, asset := range asset {
					if asset.GetMintCoin() == 0 && asset.GetOutputCoin() == 1 {
						assessments = append(assessments, Assessment{
							CourseStateToken:       string(asset.GetName()),
							CourseStateTokenOutput: output,
						})
					}
				}
			}
		}
	}

	if userToken == "" || len(assessments) == 0 {
		return nil, false
	}

	// Then Check Datum structure for the input token and the output token
	// Input datum is export type Committed = ConStr1<[ByteString, ByteString, List<ByteString>]>;
	// Output datum is export type State = ConStr0<[List<ByteString>]>;

	var results []models.Assessment

	for _, assessment := range assessments {
		if userToken[1:] != assessment.CourseStateToken {
			datum := assessment.CourseStateTokenOutput.GetDatum().GetPayload()

			constr := datum.GetConstr()
			if constr == nil {
				return nil, false
			}

			switch constr.GetTag() {
			case 121:
				if handleAccept(constr) {
					results = append(results, models.Assessment{
						StudentAlias: assessment.CourseStateToken,
						Assessment:   models.Accept,
					})
				}
			case 122:
				if handleRefuse(constr) {
					results = append(results, models.Assessment{
						StudentAlias: assessment.CourseStateToken,
						Assessment:   models.Refuse,
					})
				}
			default:
				return nil, false
			}
		}
	}

	if len(results) > 0 {
		return &models.TeacherCourseAssignmentsAssess{
			TxHash:       hex.EncodeToString(tx.GetHash()),
			Alias:        userToken[1:],
			CourseID:     "",
			AssignmentID: "",
			Assessments:  results,
		}, true
	}

	return nil, false
}

func handleAccept(constr *cardano.Constr) bool {

	// Check structure: should have 1 field that’s an Array of ByteStrings
	if len(constr.GetFields()) != 1 {
		return false
	}

	array := constr.GetFields()[0].GetArray()
	if array == nil {
		return false
	}

	// Now check each array item is a BoundedBytes
	for i, item := range array.GetItems() {
		pd := item.GetBoundedBytes()
		if pd == nil {
			fmt.Printf("Item %d is not a BoundedBytes.\n", i)
			return false
		}
	}

	return true
}

func handleRefuse(constr *cardano.Constr) bool {
	return true
}
