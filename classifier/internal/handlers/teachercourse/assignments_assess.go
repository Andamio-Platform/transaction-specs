package teachercourse

import (
	"encoding/hex"
	"fmt"
	"slices"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func AssessAssignments(tx *cardano.Tx, accessTokenPolicy string, courseStatePolicyIds []string) bool {

	type Decision string

	const (
		Accept Decision = "accept"
		Refuse Decision = "refuse"
	)

	var decision Decision = ""

	// Check outputs has one of the courseStatePolicyIds
	// Check if user (u) token is different with course state token
	var userToken string
	var courseStateToken string

	var courseStateTokenOutput *cardano.TxOutput

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
						courseStateToken = string(asset.GetName())
						courseStateTokenOutput = output
					}
				}
			}
		}
	}

	if userToken == "" || courseStateToken == "" {
		return false
	}

	// Then Check Datum structure for the input token and the output token
	// Input datum is export type Committed = ConStr1<[ByteString, ByteString, List<ByteString>]>;
	// Output datum is export type State = ConStr0<[List<ByteString>]>;

	if userToken[1:] != courseStateToken {
		datum := courseStateTokenOutput.GetDatum().GetPayload()

		constr := datum.GetConstr()
		if constr == nil {
			return false
		}

		switch constr.GetTag() {
		case 121:
			decision = Accept
			println(decision)
			return handleAccept(constr)
		case 122:
			decision = Refuse
			println(decision)
			return handleRefuse(constr)
		default:
			return false
		}

	}

	return false
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
