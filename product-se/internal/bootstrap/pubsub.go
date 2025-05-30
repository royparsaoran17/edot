package bootstrap

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"

	"product-se/internal/appctx"
	"product-se/pkg/logger"
	"product-se/pkg/pubsubx"
)

func RegistryPubSubConsumer(cfg *appctx.Config) pubsubx.Subscriberer {
	credOpt := option.WithCredentialsFile(cfg.Pubsub.AccountPath)
	cl, err := pubsub.NewClient(context.Background(), cfg.Pubsub.ProjectID, credOpt)
	if err != nil {
		logger.Fatal(fmt.Sprintf("google pusbsub conusmer error:%v", err))
	}

	return pubsubx.NewGSubscriber(cl)
}

func RegistryPubSubPublisher(cfg *appctx.Config) pubsubx.Publisher {
	credOpt := option.WithCredentialsFile(cfg.Pubsub.AccountPath)
	cl, err := pubsub.NewClient(context.Background(), cfg.Pubsub.ProjectID, credOpt)
	if err != nil {
		logger.Fatal(fmt.Sprintf("google pusbsub publisher error:%v", err))
	}

	return pubsubx.NewGPublisher(cl)
}
