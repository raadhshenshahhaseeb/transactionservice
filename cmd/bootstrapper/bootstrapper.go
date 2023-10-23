package bootstrapper

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/txservice/configuration"
	"github.com/hyperversalblocks/txservice/pkg/logger"
	"github.com/hyperversalblocks/txservice/pkg/node"
)

type Container struct {
	ctx    context.Context
	node   *node.Node
	logger *logrus.Logger
	conf   *configuration.Config
}

func New() error {
	container := new(Container)

	err := container.Init(context.Background())
	if err != nil {
		return fmt.Errorf("unable to bootstrap service")
	}

	return nil
}

func (c *Container) Init(ctx context.Context) error {
	config, err := configuration.Init()
	if err != nil {
		return fmt.Errorf("failed to load configuration file: %w", err)
	}

	logger := logger.Init(config)

	node, err := node.InitNode(c.ctx, config.Chain.PrivateKey, config.Chain.Endpoint, logger)
	if err != nil {
		return fmt.Errorf("failed to initialize node: %w", err)
	}

	c.conf = config
	c.node = node
	c.logger = logger
	c.ctx = ctx
	return nil
}
