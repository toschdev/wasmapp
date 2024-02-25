package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "wasmapp/testutil/keeper"
	"wasmapp/x/wasmapp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WasmappKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
