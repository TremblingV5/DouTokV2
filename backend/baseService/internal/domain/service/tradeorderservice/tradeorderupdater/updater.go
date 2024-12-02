package tradeorderupdater

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/tradeorder"
)

type Updater interface {
	DoUpdate(ctx context.Context, t interface{}) (*tradeorder.Entity, error)
}
