package types

import (
	"github.com/maticnetwork/tendermint/crypto"

	sdk "github.com/maticnetwork/cosmos-sdk/types"
)

// NewTestTx creates new test tx
func NewTestTx(ctx sdk.Context, msg sdk.Msg, priv crypto.PrivKey, accNum uint64, seq uint64) sdk.Tx {
	signBytes := StdSignBytes(ctx.ChainID(), accNum, seq, msg, "")
	sig, err := priv.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	tx := NewStdTx(msg, sig, "")
	return tx
}

// NewTestTxWithMemo create new test tx
func NewTestTxWithMemo(ctx sdk.Context, msg sdk.Msg, priv crypto.PrivKey, accNum uint64, seq uint64, memo string) sdk.Tx {
	signBytes := StdSignBytes(ctx.ChainID(), accNum, seq, msg, "")
	sig, err := priv.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	tx := NewStdTx(msg, sig, memo)
	return tx
}

// NewTestTxWithSignBytes creates tx with sign bytes
func NewTestTxWithSignBytes(msg sdk.Msg, priv crypto.PrivKey, accNum uint64, seq uint64, signBytes []byte, memo string) sdk.Tx {
	sig, err := priv.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	tx := NewStdTx(msg, sig, memo)
	return tx
}
