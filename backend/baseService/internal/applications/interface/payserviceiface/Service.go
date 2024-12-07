package payserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
)

type Service interface {
	CreateFeeList(ctx context.Context, request *api.CreateFeeListRequest) error
}
