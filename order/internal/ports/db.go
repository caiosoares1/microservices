package ports

import "github.com/caiosoares1/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	UpdateStatus(*domain.Order) error
}
