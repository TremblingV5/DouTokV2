package deliveryinfoapp

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/deliveryinfoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/deliveryinfo"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	ErrDeliveryInfoNotBelong2Account = "delivery info not belong to account"
)

type Application struct {
	deliveryInfoService deliveryinfoiface.Service
}

func New(deliveryInfoService deliveryinfoiface.Service) *Application {
	return &Application{
		deliveryInfoService: deliveryInfoService,
	}
}

func (a *Application) CreateDeliveryInfo(ctx context.Context, request *api.CreateDeliveryInfoRequest) (*api.CreateDeliveryInfoResponse, error) {
	info, err := a.deliveryInfoService.Create(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("CreateDeliveryInfo failed: %v", err)
		return &api.CreateDeliveryInfoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.CreateDeliveryInfoResponse{
		Meta:         utils.GetSuccessMeta(),
		DeliveryInfo: info.ToPb(),
	}, nil
}

func (a *Application) checkIsBelong2Account(ctx context.Context, accountId, deliveryInfoId int64) error {
	isBelong, err := a.deliveryInfoService.IsBelongToAccount(ctx, accountId, deliveryInfoId)
	if err != nil {
		log.Context(ctx).Errorf("UpdateDeliveryInfo failed: %v", err)
		return err
	}

	if !isBelong {
		log.Context(ctx).Warnw(
			"info", "check is belong to account not belong the given account",
			"accountId", accountId,
			"deliveryInfoId", deliveryInfoId,
		)
		return errors.New(ErrDeliveryInfoNotBelong2Account)
	}

	return nil
}

func (a *Application) UpdateDeliveryInfo(ctx context.Context, request *api.UpdateSourceDeliveryInfoRequest) (*api.UpdateSourceDeliveryInfoResponse, error) {
	err := a.checkIsBelong2Account(ctx, request.GetAccountId(), request.GetId())
	if err != nil {
		return &api.UpdateSourceDeliveryInfoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	info, err := a.deliveryInfoService.Update(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("UpdateDeliveryInfo failed: %v", err)
		return &api.UpdateSourceDeliveryInfoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.UpdateSourceDeliveryInfoResponse{
		Meta:         utils.GetSuccessMeta(),
		DeliveryInfo: info.ToPb(),
	}, nil
}

func (a *Application) QueryById(ctx context.Context, request *api.QueryByIdRequest) (*api.QueryByIdResponse, error) {
	data, err := a.deliveryInfoService.GetById(ctx, request.GetId())
	if err != nil {
		log.Context(ctx).Errorf("QueryById failed: %v", err)
		return &api.QueryByIdResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.QueryByIdResponse{
		Meta:         utils.GetSuccessMeta(),
		DeliveryInfo: data.ToPb(),
	}, nil
}

func (a *Application) QueryByAccountId(ctx context.Context, request *api.QueryByAccountIdRequest) (*api.QueryByAccountIdResponse, error) { //TODO implement me
	data, total, err := a.deliveryInfoService.GetByAccountId(
		ctx,
		request.GetAccountId(),
		request.GetPagination().GetPage(),
		request.GetPagination().GetSize(),
	)
	if err != nil {
		log.Context(ctx).Errorf("QueryByAccountId failed: %v", err)
		return &api.QueryByAccountIdResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.QueryByAccountIdResponse{
		Meta:             utils.GetSuccessMeta(),
		DeliveryInfoList: deliveryinfo.ToPbList(data),
		Pagination:       utils.GetPageResponse(total, request.GetPagination().GetPage(), request.GetPagination().GetSize()),
	}, nil
}

func (a *Application) DeleteDeliveryInfo(ctx context.Context, request *api.DeleteDeliveryInfoRequest) (*api.DeleteDeliveryInfoResponse, error) { //TODO implement me
	err := a.checkIsBelong2Account(ctx, request.GetAccountId(), request.GetId())
	if err != nil {
		return &api.DeleteDeliveryInfoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	err = a.deliveryInfoService.DeleteById(ctx, request.GetId())
	if err != nil {
		log.Context(ctx).Errorf("DeleteDeliveryInfo failed: %v", err)
		return &api.DeleteDeliveryInfoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.DeleteDeliveryInfoResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.DeliveryInfoServiceServer = (*Application)(nil)
