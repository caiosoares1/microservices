package grpc

import (
	"context"
	"fmt"

	"github.com/caiosoares1/microservices-proto/golang/shipping"
	"github.com/caiosoares1/microservices/shipping/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipping...")

	items := make([]domain.ShippingItem, len(request.Items))
	for i, item := range request.Items {
		items[i] = domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		}
	}

	newShipping := domain.NewShipping(request.OrderId, items)
	result, err := a.api.CreateShipping(ctx, newShipping)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to create shipping. %v", err)).Err()
	}

	return &shipping.CreateShippingResponse{
		ShippingId:   result.ID,
		DeliveryDays: result.DeliveryDays,
	}, nil
}
