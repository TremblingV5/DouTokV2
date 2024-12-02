package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
)

type TradeOrderRepository interface {
	QueryByIdAndBizOrderNo(ctx context.Context, idList []int64, bizOrderNoList []string) ([]*models.TradeOrder, error)
}
