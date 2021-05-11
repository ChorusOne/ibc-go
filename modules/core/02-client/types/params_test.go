package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

func TestValidateParams(t *testing.T) {
	testCases := []struct {
		name    string
		params  Params
		expPass bool
	}{
		{"default params", DefaultParams(), true},
		{"custom params", NewParams(true, exported.Tendermint), true},
		{"blank client", NewParams(false, " "), false},
	}

	for _, tc := range testCases {
		err := tc.params.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
		}
	}
}
