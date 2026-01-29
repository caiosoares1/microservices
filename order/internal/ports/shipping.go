package ports

import "github.com/caiosoares1/microservices/order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipping(order *domain.Order) (int32, error)
}
