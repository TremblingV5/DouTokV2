package deliveryinfoappproviders

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/deliveryinfoapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/deliveryinfoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/deliveryinfoservice"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/deliveryinforepo"
	"github.com/google/wire"
)

var DeliveryInfoRepoProviderSet = wire.NewSet(
	deliveryinforepo.New,
	wire.Bind(new(repoiface.DeliveryInfoRepository), new(*deliveryinforepo.Repository)),
)

var DeliveryInfoServiceProviderSet = wire.NewSet(
	deliveryinfoservice.New,
	wire.Bind(new(deliveryinfoiface.Service), new(*deliveryinfoservice.Service)),
)

var DeliveryInfoAppProviderSet = wire.NewSet(
	deliveryinfoapp.New,
	DeliveryInfoRepoProviderSet,
	DeliveryInfoServiceProviderSet,
)
