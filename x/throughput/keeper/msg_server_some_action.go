package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"throughput/x/throughput/types"
)

func (k msgServer) SomeAction(goCtx context.Context, msg *types.MsgSomeAction) (*types.MsgSomeActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSomeActionResponse{}, nil
}
