package main

import (
	"context"
	"encoding/hex"

	"connectrpc.com/connect"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/query"
	sdk "github.com/utxorpc/go-sdk"
)

func main() {

	hashHex := "b5eaffde5f818310567881b7c14d9e071a29b0c20cefbc73cec3f350da9aac3d"
	tx := GetCardanoTx(hashHex)

	TeacherCourseModulesManage(tx)
}

func GetCardanoTx(hashHex string) *cardano.Tx {
	baseUrl := "https://preprod.utxorpc.dolos.andamio.space"
	ctx := context.Background()

	client := sdk.NewQueryServiceClient(sdk.NewClient(sdk.WithBaseUrl(baseUrl)))

	hash, err := hex.DecodeString(hashHex)
	if err != nil {
		sdk.HandleError(err)
	}

	searchRequest := connect.NewRequest(&query.ReadTxRequest{
		Hash: hash,
	})

	resp, err := client.ReadTx(ctx, searchRequest)
	if err != nil {
		sdk.HandleError(err)
	}

	return resp.Msg.Tx.GetCardano()
}
