package models

import (
	"fmt"
	orderService "github.com/AleksK1NG/es-microservice/proto/order"
	"time"
)

type Order struct {
	ID              string      `json:"id" bson:"_id,omitempty"`
	ShopItems       []*ShopItem `json:"shopItems" bson:"shopItems,omitempty"`
	AccountEmail    string      `json:"accountEmail" bson:"accountEmail,omitempty"`
	DeliveryAddress string      `json:"deliveryAddress" bson:"deliveryAddress,omitempty"`
	CancelReason    string      `json:"cancelReason" bson:"cancelReason,omitempty"`
	TotalPrice      float64     `json:"totalPrice" bson:"totalPrice,omitempty"`
	DeliveredTime   time.Time   `json:"deliveredTime" bson:"deliveredTime,omitempty"`
	Created         bool        `json:"created" bson:"created,omitempty"`
	Paid            bool        `json:"paid" bson:"paid,omitempty"`
	Submitted       bool        `json:"submitted" bson:"submitted,omitempty"`
	Delivering      bool        `json:"delivering" bson:"delivering,omitempty"`
	Delivered       bool        `json:"delivered" bson:"delivered,omitempty"`
	Canceled        bool        `json:"canceled" bson:"canceled,omitempty"`
}

func (o *Order) String() string {
	return fmt.Sprintf("ID: {%s}, ShopItems: {%+v}, Created: {%v}, Paid: {%v}, Submitted: {%v}, Delivering: {%v}, Delivered: {%v}, Canceled: {%v}, TotalPrice: {%v}, AccountEmail: {%s},",
		o.ID,
		o.ShopItems,
		o.Created,
		o.Paid,
		o.Submitted,
		o.Delivering,
		o.Delivered,
		o.Canceled,
		o.TotalPrice,
		o.AccountEmail,
	)
}

func NewOrder() *Order {
	return &Order{
		ShopItems:  make([]*ShopItem, 0),
		Created:    false,
		Paid:       false,
		Submitted:  false,
		Delivering: false,
		Delivered:  false,
		Canceled:   false,
	}
}

func OrderToProto(order *Order, id string) *orderService.Order {
	return &orderService.Order{
		ID:         id,
		ShopItems:  ShopItemsToProto(order.ShopItems),
		Created:    order.Created,
		Paid:       order.Paid,
		Submitted:  order.Submitted,
		Delivering: order.Delivering,
		Delivered:  order.Delivered,
		Canceled:   order.Canceled,
	}
}
