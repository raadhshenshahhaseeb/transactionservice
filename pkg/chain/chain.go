package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func Init(rpcEndpoint string, ctx context.Context) (*rpc.Client, error) {
	rpcClient, err := rpc.DialContext(ctx, rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %w", err)
	}

	var versionString string
	err = rpcClient.CallContext(ctx, &versionString, "web3_clientVersion")
	if err != nil {
		return nil, fmt.Errorf("eth client get version: %w", err)
	}

	return rpcClient, nil
}
