package server

import (
	"context"
	"fmt"
	"github.com/hyperversalblocks/txservice/pkg/transaction"
	"time"
)

func (s *Server) Init(ctx context.Context) error {
	err := s.Init(ctx)
	if err != nil {
		return fmt.Errorf("error init chain: %w", err)
	}
	// Sync the with the given Ethereum backend:
	isSynced, _, err := transaction.IsSynced(ctx, s.Backend, 1*time.Minute)
	if err != nil {
		return fmt.Errorf("is synced: %w", err)
	}
	if !isSynced {
		s.logger.Info("waiting to sync with the Ethereum backend")

		err := transaction.WaitSynced(ctx, s.Backend, 1*time.Minute)
		if err != nil {
			return fmt.Errorf("waiting backend sync: %w", err)
		}
	}

	return nil
}
