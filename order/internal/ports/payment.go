package ports

import "github.com/caiosoares1/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order domain.Order) error
}
