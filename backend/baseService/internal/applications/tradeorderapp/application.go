package tradeorderapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/tradeorderservice"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/tradeorder"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	tradeOrderService tradeorderservice.TradeOrderService
}

func New(
	tradeOrderService tradeorderservice.TradeOrderService,
) *Application {
	return &Application{
		tradeOrderService: tradeOrderService,
	}
}

func (t *Application) QueryTradeOrder(ctx context.Context, request *api.QueryTradeOrderRequest) (*api.QueryTradeOrderResponse, error) {
	tradeOrderList, err := t.tradeOrderService.QueryTradeOrder(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("query trade order failed: %v", err)
		return &api.QueryTradeOrderResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.QueryTradeOrderResponse{
		Meta:           utils.GetSuccessMeta(),
		TradeOrderList: tradeorder.ToPb(tradeOrderList),
	}, nil
}

func (t *Application) TradeCreate(ctx context.Context, request *api.TradeCreateRequest) (*api.TradeCreateResponse, error) {
	tradeOrder, err := t.tradeOrderService.Create(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("create trade order failed: %v", err)
		return &api.TradeCreateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.TradeCreateResponse{
		Meta:       utils.GetSuccessMeta(),
		TradeOrder: tradeOrder.ToPB(),
	}, nil
}

func (t *Application) MergeTrade(ctx context.Context, request *api.MergeTradeRequest) (*api.MergeTradeResponse, error) {
	tradeOrder, err := t.tradeOrderService.Merge(ctx, request.GetOriginalOrderId(), request.GetOrderIdList()...)
	if err != nil {
		log.Context(ctx).Errorf("merge trade order failed: %v", err)
		return &api.MergeTradeResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.MergeTradeResponse{
		Meta:       utils.GetSuccessMeta(),
		TradeOrder: tradeOrder.ToPB(),
	}, nil
}

func (t *Application) UpdateExtendInfo(ctx context.Context, request *api.UpdateExtendInfoRequest) (*api.UpdateExtendInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Application) UpdateTradeOrderStatus(ctx context.Context, request *api.UpdateTradeOrderStatusRequest) (*api.UpdateTradeOrderStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Application) UpdateDeliveryInfo(ctx context.Context, request *api.UpdateDeliveryInfoRequest) (*api.UpdateDeliveryInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Application) AddSubOrder(ctx context.Context, request *api.AddSubOrderRequest) (*api.AddSubOrderResponse, error) {
	tradeOrder, err := t.tradeOrderService.AddSubOrder(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("add sub order failed: %v", err)
		return &api.AddSubOrderResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.AddSubOrderResponse{
		Meta:       utils.GetSuccessMeta(),
		TradeOrder: tradeOrder.ToPB(),
	}, nil
}

var _ api.TradeServiceServer = (*Application)(nil)
