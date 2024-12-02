package deliveryinfoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/deliveryinfo"
)

type Service interface {
	Create(ctx context.Context, in *api.CreateDeliveryInfoRequest) (*deliveryinfo.Entity, error)
	IsBelongToAccount(ctx context.Context, accountId int64, deliveryInfoId int64) (bool, error)
	Update(ctx context.Context, in *api.UpdateSourceDeliveryInfoRequest) (*deliveryinfo.Entity, error)
	GetById(ctx context.Context, id int64) (*deliveryinfo.Entity, error)
	GetByAccountId(ctx context.Context, accountId int64, page, size int32) ([]*deliveryinfo.Entity, int64, error)
	DeleteById(ctx context.Context, id int64) error
}
