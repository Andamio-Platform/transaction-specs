package teachercourse

import "github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"

func AssessAssignments(tx *cardano.Tx, courseStatePolicyIds []string) bool {
	// Check outputs has one of the courseStatePolicyIds

	// Check if user (u) token is different with course state token

	// Then Check Datum structure for the input token and the output token
	// Input datum is export type Committed = ConStr1<[ByteString, ByteString, List<ByteString>]>;
	// Output datum is export type State = ConStr0<[List<ByteString>]>;

	return false
}
