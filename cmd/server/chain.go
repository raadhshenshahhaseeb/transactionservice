package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/hyperversalblocks/txservice/pkg/transaction"
)

type Chain struct {
	SenderAddress common.Address
	PrivateKey    *ecdsa.PrivateKey
	ChainID       int64
}

func (s *Server) InitChain(ctx context.Context) error {
	rpcClient, err := rpc.DialContext(ctx, s.config.Chain.RPCEndpoint)
	if err != nil {
		return fmt.Errorf("dial eth client: %w", err)
	}

	var versionString string
	err = rpcClient.CallContext(ctx, &versionString, "web3_clientVersion")
	if err != nil {
		s.logger.Info("could not connect to backend; check your endpoint config var", "backend_endpoint", s.config.Chain.RPCEndpoint)
		return fmt.Errorf("eth client get version: %w", err)
	}

	backend := transaction.NewBackend(ethclient.NewClient(rpcClient))

	s.logger.Info("connected to ethereum backend", "version", versionString)

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("get chain id: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(s.config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Public Key Error")
	}

	publicKeyECDSA, ok = publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Public Key Error")
	}

	SenderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	s.chain = &Chain{
		SenderAddress: SenderAddress,
		PrivateKey:    privateKey,
		ChainID:       chainID.Int64(),
	}

	s.Backend = backend

	return nil
}
