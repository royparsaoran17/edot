package order

import (
	"fmt"
	"manage-se/internal/appctx"
	"manage-se/internal/provider/dependencies"
	"strings"
)

type client struct {
	cfg *appctx.Order
	dep *dependencies.Dependency
}

func NewClient(cfg *appctx.Order, dependency *dependencies.Dependency) *client {
	return &client{
		cfg: cfg,
		dep: dependency,
	}
}

func (c *client) endpoint(path string) string {
	return fmt.Sprintf("%s/%s", strings.TrimSuffix(c.cfg.BaseURL, "/"), strings.TrimPrefix(path, "/"))
}
