package tradeorderservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/tradeorder"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/tradeorderservice/tradeorderupdater"
)

type TradeOrderService interface {
	QueryTradeOrder(ctx context.Context, in *api.QueryTradeOrderRequest) ([]*tradeorder.Entity, error)
	Create(ctx context.Context, in *api.TradeCreateRequest) (*tradeorder.Entity, error)
	Merge(ctx context.Context, originalId int64, subOrderIdList ...int64) (*tradeorder.Entity, error)
	AddSubOrder(ctx context.Context, in *api.AddSubOrderRequest) (*tradeorder.Entity, error)
	Update(ctx context.Context, updater tradeorderupdater.Updater) (*tradeorder.Entity, error)
}
