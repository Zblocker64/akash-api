package client

import (
	"context"
	"errors"
	"strings"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/pflag"
	tmjclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"

	"github.com/akash-network/akash-api/go/node/client/v1beta2"
)

var (
	ErrUnknownClientVersion = errors.New("akash-api: unknown client version")
)

func DiscoverClient(ctx context.Context, cctx sdkclient.Context, flags *pflag.FlagSet, setup func(interface{}) error) error {
	rpc, err := tmjclient.New(cctx.NodeURI)
	if err != nil {
		return err
	}

	result := new(Akash)
	params := make(map[string]interface{})
	_, err = rpc.Call(ctx, "akash", params, result)
	if err != nil && !strings.Contains(err.Error(), "Method not found") {
		return err
	}

	// if client info is nil, mostly likely "akash" endpoint is not yet supported on the node
	// fallback to manually set version to v1beta2
	if result.ClientInfo == nil {
		result.ClientInfo = &ClientInfo{ApiVersion: "v1beta2"}
	}

	var cl interface{}

	switch result.ClientInfo.ApiVersion {
	case "v1beta2":
		cl, err = v1beta2.NewClient(ctx, cctx, flags)
	default:
		return ErrUnknownClientVersion
	}

	if err = setup(cl); err != nil {
		return err
	}

	return nil
}

func DiscoverQueryClient(ctx context.Context, cctx sdkclient.Context, setup func(interface{}) error) error {
	rpc, err := tmjclient.New(cctx.NodeURI)
	if err != nil {
		return err
	}

	result := new(Akash)
	params := make(map[string]interface{})
	_, err = rpc.Call(ctx, "akash", params, result)
	if err != nil && !strings.Contains(err.Error(), "Method not found") {
		return err
	}

	if result.ClientInfo == nil {
		result.ClientInfo = &ClientInfo{ApiVersion: "v1beta2"}
	}

	var cl interface{}

	switch result.ClientInfo.ApiVersion {
	case "v1beta2":
		cl = v1beta2.NewQueryClient(cctx)
	default:
		return ErrUnknownClientVersion
	}

	if err = setup(cl); err != nil {
		return err
	}

	return nil
}
