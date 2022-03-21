package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/axelarnetwork/axelar-core/x/evm/types"
)

var _ types.QueryServiceServer = Querier{}

// Querier implements the grpc querier
type Querier struct {
	keeper types.BaseKeeper
	nexus  types.Nexus
}

// NewGRPCQuerier returns a new Querier
func NewGRPCQuerier(k types.BaseKeeper, n types.Nexus) Querier {
	return Querier{
		keeper: k,
		nexus:  n,
	}
}

// BurnerInfo implements the burner info grpc query
func (q Querier) BurnerInfo(c context.Context, req *types.BurnerInfoRequest) (*types.BurnerInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chains := getEVMChains(ctx, q.nexus)
	if len(chains) == 0 {
		return nil, status.Error(codes.NotFound, "no chains registered")
	}

	for _, chain := range chains {
		ck := q.keeper.ForChain(chain)
		burnerInfo := ck.GetBurnerInfo(ctx, req.Address)
		if burnerInfo != nil {
			return &types.BurnerInfoResponse{Chain: ck.GetParams(ctx).Chain, BurnerInfo: burnerInfo}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "unknown address")
}

// ConfirmationHeight implements the confirmation height grpc query
func (q Querier) ConfirmationHeight(c context.Context, req *types.ConfirmationHeightRequest) (*types.ConfirmationHeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, ok := q.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, status.Error(codes.NotFound, "unknown chain")

	}

	ck := q.keeper.ForChain(string(req.Chain))
	height, ok := ck.GetRequiredConfirmationHeight(ctx)
	if !ok {
		return nil, status.Error(codes.NotFound, "could not get confirmation height")
	}

	return &types.ConfirmationHeightResponse{Height: height}, nil
}

// DepositState fetches the state of a deposit confirmation using a grpc query
func (q Querier) DepositState(c context.Context, req *types.DepositStateRequest) (*types.DepositStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	ck := q.keeper.ForChain(req.Chain)

	s, log, code := queryDepositState(ctx, ck, q.nexus, req.Params)
	if code != codes.OK {
		return nil, status.Error(code, log)
	}

	return &types.DepositStateResponse{Status: s}, nil
}

// PendingCommands returns the pending commands from a gateway
func (q Querier) PendingCommands(c context.Context, req *types.PendingCommandsRequest) (*types.PendingCommandsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, ok := q.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("chain %s not found", req.Chain))
	}

	ck := q.keeper.ForChain(req.Chain)

	var commands []types.QueryCommandResponse
	for _, cmd := range ck.GetPendingCommands(ctx) {
		cmdResp, err := GetCommandResponse(ctx, ck.GetName(), q.nexus, cmd)
		if err != nil {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		commands = append(commands, cmdResp)
	}

	return &types.PendingCommandsResponse{Commands: commands}, nil
}
