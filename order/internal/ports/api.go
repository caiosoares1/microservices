package ports

import "github.com/caiosoares1/microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}