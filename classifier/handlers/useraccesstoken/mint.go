package useraccesstoken

import (
	"encoding/hex"

	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/models"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"
)

func Mint(tx *cardano.Tx) (*models.UserAccessTokenMint, bool) {
	accessTokenPolicy := config.Get().CurrentV2().IndexMS.MSCPolicyID
	mints := tx.GetMint()

	if len(mints) > 0 {
		for _, mint := range mints {
			for _, asset := range mint.GetAssets() {
				if asset.MintCoin > 0 {
					if hex.EncodeToString(mint.GetPolicyId()) == accessTokenPolicy {
						redeemer := mint.GetRedeemer().GetPayload()
						alias := string(redeemer.GetBoundedBytes())
						return &models.UserAccessTokenMint{
							TxHash: hex.EncodeToString(tx.GetHash()),
							Alias:  alias,
						}, true
					}
				}
			}
		}

	}
	return nil, false
}
