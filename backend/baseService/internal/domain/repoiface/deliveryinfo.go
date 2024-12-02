package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/deliveryinfo"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
)

type DeliveryInfoRepository interface {
	Create(ctx context.Context, deliveryInfo *models.DeliveryInfo) error
	IsBelong2Account(ctx context.Context, accountId, deliveryInfoId int64) (bool, error)
	UpdateById(ctx context.Context, deliveryInfoId int64, deliveryInfo *deliveryinfo.Entity) error
	GetById(ctx context.Context, id int64) (*models.DeliveryInfo, error)
	GetByAccountId(ctx context.Context, accountId int64, page, size int32) ([]*models.DeliveryInfo, error)
	CountByAccountId(ctx context.Context, accountId int64) (int64, error)
	DeleteById(ctx context.Context, id int64) error
}
