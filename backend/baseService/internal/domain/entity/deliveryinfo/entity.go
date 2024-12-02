package deliveryinfo

import (
	"encoding/json"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"time"
)

type Entity struct {
	ID              int64
	AccountID       int64
	ReceiverName    string
	ReceiverPhone   string
	ReceiverAddress string
	Extra           map[string]string
	IsDeleted       bool
	CreateTime      time.Time
	UpdateTime      time.Time
}

func NewWithModel(model *models.DeliveryInfo) *Entity {
	extra := make(map[string]string)
	_ = json.Unmarshal([]byte(model.Extra), &extra)
	return &Entity{
		ID:              model.ID,
		AccountID:       model.AccountID,
		ReceiverName:    model.ReceiverName,
		ReceiverPhone:   model.ReceiverPhone,
		ReceiverAddress: model.ReceiverAddress,
		Extra:           extra,
		IsDeleted:       model.IsDeleted,
		CreateTime:      model.CreateTime,
		UpdateTime:      model.UpdateTime,
	}
}

func NewWithModelList(modelList []*models.DeliveryInfo) []*Entity {
	var entityList []*Entity
	for _, model := range modelList {
		entityList = append(entityList, NewWithModel(model))
	}
	return entityList
}

func NewWithUpdateReq(in *api.UpdateSourceDeliveryInfoRequest) *Entity {
	return &Entity{
		AccountID:       in.AccountId,
		ReceiverName:    in.ReceiverName,
		ReceiverPhone:   in.ReceiverPhone,
		ReceiverAddress: in.ReceiverAddress,
	}
}

func (e *Entity) ToPb() *api.SourceDeliveryInfo {
	return &api.SourceDeliveryInfo{
		Id:              e.ID,
		AccountId:       e.AccountID,
		ReceiverName:    e.ReceiverName,
		ReceiverPhone:   e.ReceiverPhone,
		ReceiverAddress: e.ReceiverAddress,
		Extra:           e.Extra,
	}
}

func ToPbList(in []*Entity) []*api.SourceDeliveryInfo {
	if in == nil {
		return nil
	}

	var out []*api.SourceDeliveryInfo
	for _, e := range in {
		out = append(out, e.ToPb())
	}
	return out
}
