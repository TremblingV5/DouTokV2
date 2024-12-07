package payapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/payserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	payService payserviceiface.Service
}

func New(payService payserviceiface.Service) *Application {
	return &Application{
		payService: payService,
	}
}

func (a *Application) CreateFeeList(ctx context.Context, request *api.CreateFeeListRequest) (*api.CreateFeeListResponse, error) {
	err := a.payService.CreateFeeList(ctx, request)
	if err != nil {
		log.Context(ctx).Errorf("create fee list failed: %v", err)
		return &api.CreateFeeListResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.CreateFeeListResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) Pay(ctx context.Context, request *api.PayRequest) (*api.PayResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ api.PayServiceServer = (*Application)(nil)
