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

	hashHex := "64b0b7fe30a6e34ade9bd489e1bdc72ef5495f7c56e9b30154851cf4812a06cc"
	tx := GetCardanoTx(hashHex)

	_ = hex.EncodeToString(tx.Hash)
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
