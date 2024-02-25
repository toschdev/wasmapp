package wasmapp_test

import (
	"testing"

	keepertest "wasmapp/testutil/keeper"
	"wasmapp/testutil/nullify"
	wasmapp "wasmapp/x/wasmapp/module"
	"wasmapp/x/wasmapp/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WasmappKeeper(t)
	wasmapp.InitGenesis(ctx, k, genesisState)
	got := wasmapp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
