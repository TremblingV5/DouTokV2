package deliveryinforepo

import (
	"context"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/deliveryinfo"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, deliveryInfo *models.DeliveryInfo) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		return tx.WithContext(ctx).DeliveryInfo.Create(deliveryInfo)
	})
}

func (r *Repository) IsBelong2Account(ctx context.Context, accountId, deliveryInfoId int64) (bool, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (bool, error) {
		count, err := tx.WithContext(ctx).DeliveryInfo.Where(
			query.DeliveryInfo.AccountID.Eq(accountId),
			query.DeliveryInfo.ID.Eq(deliveryInfoId),
		).Count()
		if err != nil {
			return false, err
		}

		return count > 0, nil
	})
}

func (r *Repository) UpdateById(ctx context.Context, deliveryInfoId int64, deliveryInfo *deliveryinfo.Entity) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := tx.WithContext(ctx).DeliveryInfo.Updates(deliveryInfo)
		return err
	})
}

func (r *Repository) GetById(ctx context.Context, id int64) (*models.DeliveryInfo, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (*models.DeliveryInfo, error) {
		return tx.WithContext(ctx).DeliveryInfo.Where(
			query.DeliveryInfo.ID.Eq(id),
		).First()
	})
}

func (r *Repository) GetByAccountId(ctx context.Context, accountId int64, page, size int32) ([]*models.DeliveryInfo, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) ([]*models.DeliveryInfo, error) {
		return tx.WithContext(ctx).DeliveryInfo.Where(
			query.DeliveryInfo.AccountID.Eq(accountId),
		).Limit(int(size)).Offset(int(page * size)).Find()
	})
}

func (r *Repository) CountByAccountId(ctx context.Context, accountId int64) (int64, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (int64, error) {
		return tx.WithContext(ctx).DeliveryInfo.Where(
			query.DeliveryInfo.AccountID.Eq(accountId),
		).Count()
	})
}

func (r *Repository) DeleteById(ctx context.Context, id int64) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := tx.WithContext(ctx).DeliveryInfo.Where(
			query.DeliveryInfo.ID.Eq(id),
		).Delete()
		return err
	})
}

var _ repoiface.DeliveryInfoRepository = (*Repository)(nil)
