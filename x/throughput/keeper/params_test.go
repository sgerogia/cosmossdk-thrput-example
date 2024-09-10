package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "throughput/testutil/keeper"
	"throughput/x/throughput/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ThroughputKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
