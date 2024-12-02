package tradeorder

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"math/big"
	"time"
)

type Entity struct {
	ID                     int64
	BizType                api.TradeBizType
	SaleWay                api.TradeSaleWay
	PaySequence            api.TradePaySequence
	GoodsType              api.TradeGoodsType
	Status                 api.TradeOrderStatusEnum
	BizStatus              int64
	BizOrderNo             string
	BizId                  int64
	BuyerId                int64
	SellerId               int64
	PayType                api.PayType
	PayableAmount          *big.Float
	PaidAmount             *big.Float
	PromotionAmount        *big.Float
	DeliveryInfoId         int64 // TODO: delivery info
	DeliveryTime           *time.Time
	PaymentInfoID          int64
	PayTime                *time.Time
	CloseTime              *time.Time
	CloseType              api.CloseType
	CloseReason            string
	Extra                  map[string]string
	TradeOrderRelation     api.TradeOrderRelation
	OriginOrderId          int64
	ExpectPayTimeout       *time.Time
	ActualPayTimeout       *time.Time
	ExecuteAfterPayTimeout bool
	IsDeleted              bool
	CreateTime             time.Time
	UpdateTime             time.Time

	SubOrders []*Entity
}

func (e *Entity) ToPB() *api.TradeOrder {
	subOrders := make([]*api.TradeOrder, 0)
	for _, subOrder := range e.SubOrders {
		subOrders = append(subOrders, subOrder.ToPB())
	}

	return &api.TradeOrder{
		OrderId:     e.ID,
		BizType:     e.BizType,
		SaleWay:     e.SaleWay,
		PaySequence: e.PaySequence,
		GoodsType:   e.GoodsType,
		OrderStatus: e.Status,
		BizStatus:   e.BizStatus,
		BizOrderNo:  e.BizOrderNo,
		BizId:       e.BizId,
		Buyer: &api.TradeParticipant{
			AccountId: e.BuyerId,
		},
		Seller: &api.TradeParticipant{
			AccountId: e.SellerId,
		},
		PayType:         e.PayType,
		PayableAmount:   utils.Money2Ui64(e.PayableAmount),
		PaidAmount:      utils.Money2Ui64(e.PaidAmount),
		PromotionAmount: utils.Money2Ui64(e.PromotionAmount),
		//Promotions: []*api.Promotion{}{},
		DeliveryInfo: &api.DeliveryInfo{},
		CreateTime:   e.CreateTime.UnixMilli(),
		PayTime:      e.PayTime.UnixMilli(),
		DeliveryTime: e.DeliveryTime.UnixMilli(),
		CloseTime:    e.CloseTime.UnixMilli(),
		TimeoutInfo:  &api.TimeoutInfo{},
		CloseType:    e.CloseType,
		CloseReason:  e.CloseReason,
		SubOrder:     subOrders,
		//Payments: []*api.PaymentInfo{}{},
		Extra:              e.Extra,
		TradeOrderRelation: e.TradeOrderRelation,
		OriginOrderId:      e.OriginOrderId,
		//Refunds:
	}
}

func ToPb(e []*Entity) []*api.TradeOrder {
	pb := make([]*api.TradeOrder, 0)
	for _, entity := range e {
		pb = append(pb, entity.ToPB())
	}
	return pb
}

func NewWithModel(model *models.TradeOrder) *Entity {
	return &Entity{}
}

func NewListWithModels(models []*models.TradeOrder) []*Entity {
	entities := make([]*Entity, 0)
	for _, model := range models {
		entities = append(entities, NewWithModel(model))
	}
	return entities
}
