package domain

import (
	"time"
)

type ShippingItem struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type Shipping struct {
	ID           int64          `json:"id"`
	OrderID      int64          `json:"order_id"`
	Status       string         `json:"status"`
	DeliveryDays int32          `json:"delivery_days"`
	Items        []ShippingItem `json:"items"`
	CreatedAt    int64          `json:"created_at"`
}

func NewShipping(orderId int64, items []ShippingItem) Shipping {
	return Shipping{
		CreatedAt: time.Now().Unix(),
		Status:    "Pending",
		OrderID:   orderId,
		Items:     items,
	}
}

func (s *Shipping) CalculateDeliveryDays() int32 {
	var totalUnits int32
	for _, item := range s.Items {
		totalUnits += item.Quantity
	}

	deliveryDays := int32(1) + (totalUnits / 5)
	return deliveryDays
}
