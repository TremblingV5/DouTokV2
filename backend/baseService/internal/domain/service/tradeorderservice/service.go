package tradeorderservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/tradeorderservice"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/tradeorder"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/tradeorderservice/tradeorderupdater"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	tradeOrderRepo repoiface.TradeOrderRepository
}

func New(tradeOrderRepo repoiface.TradeOrderRepository) *Service {
	return &Service{
		tradeOrderRepo: tradeOrderRepo,
	}
}

func (s *Service) QueryTradeOrder(ctx context.Context, in *api.QueryTradeOrderRequest) ([]*tradeorder.Entity, error) {
	doList, err := s.tradeOrderRepo.QueryByIdAndBizOrderNo(ctx, in.OrderIdList, in.BizOrderNoList)
	if err != nil {
		log.Context(ctx).Errorf("query trade order failed: %v", err)
		return nil, err
	}

	return tradeorder.NewListWithModels(doList), nil
}

func (s *Service) Create(ctx context.Context, in *api.TradeCreateRequest) (*tradeorder.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Merge(ctx context.Context, originalId int64, subOrderIdList ...int64) (*tradeorder.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddSubOrder(ctx context.Context, in *api.AddSubOrderRequest) (*tradeorder.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, updater tradeorderupdater.Updater) (*tradeorder.Entity, error) {
	//TODO implement me
	panic("implement me")
}

var _ tradeorderservice.TradeOrderService = (*Service)(nil)
