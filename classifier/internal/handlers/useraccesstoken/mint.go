package useraccesstoken

import (
	"encoding/hex"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/cardano"

	"github.com/andamio-platform/transaction-specs/classifier/internal/models"
)

func Mint(tx *cardano.Tx, accessTokenPolicy string) (*models.UserAccessTokenMint, bool) {
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
