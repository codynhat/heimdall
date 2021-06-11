package types

import (
	"testing"

	sdk "github.com/maticnetwork/cosmos-sdk/types"
	assert "github.com/stretchr/testify/require"
)

func TestGetPulpHash(t *testing.T) {
	tc := struct {
		in  sdk.Msg
		out []byte
	}{
		in:  sdk.NewTestMsg(nil),
		out: []byte{142, 88, 179, 79},
	}
	out := GetPulpHash(tc.in)
	assert.Equal(t, string(tc.out), string(out))
}
