package api

import (
	"context"

	"github.com/caiosoares1/microservices/shipping/internal/application/core/domain"
	"github.com/caiosoares1/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) CreateShipping(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error) {
	shipping.DeliveryDays = shipping.CalculateDeliveryDays()
	shipping.Status = "Processing"

	err := a.db.Save(ctx, &shipping)
	if err != nil {
		return domain.Shipping{}, err
	}

	return shipping, nil
}
