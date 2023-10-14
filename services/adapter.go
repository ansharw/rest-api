package services

import (
	"context"
	"time"

	"github.com/ansharw/rest-api/api"
)

type ISession interface {
	CreateSession(ctx context.Context, table, key string, data *api.Session, ttl time.Duration) error
}