package throughput_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "throughput/testutil/keeper"
	"throughput/testutil/nullify"
	"throughput/x/throughput/module"
	"throughput/x/throughput/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ThroughputKeeper(t)
	throughput.InitGenesis(ctx, k, genesisState)
	got := throughput.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
