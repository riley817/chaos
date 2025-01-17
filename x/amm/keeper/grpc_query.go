package keeper

import (
	"context"
	"github.com/cosmos-builders/chaos/x/amm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{k}
}

func (k Querier) Params(c context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Querier) Pairs(c context.Context, request *types.QueryPairsRequest) (*types.QueryPairsResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	planStore := prefix.NewStore(store, types.PairKeyPrefix)
	var pairs []types.Pair
	pageRes, err := query.Paginate(planStore, request.Pagination, func(key, value []byte) error {
		var pair types.Pair
		if err := k.cdc.Unmarshal(value, &pair); err != nil {
			return err
		}
		pairs = append(pairs, pair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPairsResponse{
		Pairs:      pairs,
		Pagination: pageRes,
	}, nil
}

func (k Querier) Pair(c context.Context, request *types.QueryPairRequest) (*types.QueryPairResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	pair, found := k.GetPair(ctx, request.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "pair not found")
	}
	return &types.QueryPairResponse{Pair: pair}, nil
}
