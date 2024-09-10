package keeper

import (
	"throughput/x/throughput/types"
)

var _ types.QueryServer = Keeper{}
