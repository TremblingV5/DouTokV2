package deliveryinfoservice

import (
	"context"
	"encoding/json"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/deliveryinfoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/deliveryinfo"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	deliveryInfoRepo repoiface.DeliveryInfoRepository
}

func New(deliveryInfoRepo repoiface.DeliveryInfoRepository) *Service {
	return &Service{
		deliveryInfoRepo: deliveryInfoRepo,
	}
}

func (s *Service) newModelWithRequest(in *api.CreateDeliveryInfoRequest) (*models.DeliveryInfo, error) {
	extra, err := json.Marshal(in.Extra)
	if err != nil {
		return nil, err
	}

	return &models.DeliveryInfo{
		AccountID:       in.AccountId,
		ReceiverName:    in.ReceiverName,
		ReceiverPhone:   in.ReceiverPhone,
		ReceiverAddress: in.ReceiverAddress,
		Extra:           string(extra),
	}, nil
}

func (s *Service) Create(ctx context.Context, in *api.CreateDeliveryInfoRequest) (entity *deliveryinfo.Entity, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	model, err := s.newModelWithRequest(in)
	if err != nil {
		return nil, err
	}

	if err := s.deliveryInfoRepo.Create(ctx, model); err != nil {
		return nil, err
	}

	return deliveryinfo.NewWithModel(model), nil
}

func (s *Service) IsBelongToAccount(ctx context.Context, accountId int64, deliveryInfoId int64) (isBelong bool, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	return s.deliveryInfoRepo.IsBelong2Account(ctx, accountId, deliveryInfoId)
}

func (s *Service) Update(ctx context.Context, in *api.UpdateSourceDeliveryInfoRequest) (entity *deliveryinfo.Entity, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	entity = deliveryinfo.NewWithUpdateReq(in)
	err = s.deliveryInfoRepo.UpdateById(ctx, in.GetId(), entity)
	if err != nil {
		log.Context(ctx).Errorf("UpdateById failed: %v", err)
		return nil, err
	}

	return entity, nil
}

func (s *Service) GetById(ctx context.Context, id int64) (entity *deliveryinfo.Entity, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	model, err := s.deliveryInfoRepo.GetById(ctx, id)
	if err != nil {
		log.Context(ctx).Errorf("GetById failed: %v", err)
		return nil, err
	}

	return deliveryinfo.NewWithModel(model), nil
}

func (s *Service) GetByAccountId(ctx context.Context, accountId int64, page, size int32) (data []*deliveryinfo.Entity, total int64, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	modelList, err := s.deliveryInfoRepo.GetByAccountId(ctx, accountId, page, size)
	if err != nil {
		log.Context(ctx).Errorf("GetByAccountId failed: %v", err)
		return nil, 0, err
	}

	total, err = s.deliveryInfoRepo.CountByAccountId(ctx, accountId)
	if err != nil {
		log.Context(ctx).Errorf("CountByAccountId failed: %v", err)
		return nil, 0, err
	}

	entities := deliveryinfo.NewWithModelList(modelList)
	return entities, total, nil
}

func (s *Service) DeleteById(ctx context.Context, id int64) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	err = s.deliveryInfoRepo.DeleteById(ctx, id)
	return err
}

var _ deliveryinfoiface.Service = (*Service)(nil)
