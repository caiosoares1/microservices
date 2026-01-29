package api

import (
	"log"

	"github.com/caiosoares1/microservices/order/internal/application/core/domain"
	"github.com/caiosoares1/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	var totalItems int32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}
	if totalItems > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Order with more than 50 items is not allowed.")
	}

	var productCodes []string
	for _, item := range order.OrderItems {
		productCodes = append(productCodes, item.ProductCode)
	}

	stockItems, err := a.db.GetStockItems(productCodes)
	if err != nil {
		return domain.Order{}, status.Errorf(codes.Internal, "Failed to verify stock items: %v", err)
	}

	stockMap := make(map[string]domain.StockItem)
	for _, stockItem := range stockItems {
		stockMap[stockItem.ProductCode] = stockItem
	}

	var missingItems []string
	for _, item := range order.OrderItems {
		if _, exists := stockMap[item.ProductCode]; !exists {
			missingItems = append(missingItems, item.ProductCode)
		}
	}

	if len(missingItems) > 0 {
		return domain.Order{}, status.Errorf(codes.NotFound, "Items not found in stock: %v", missingItems)
	}

	err = a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		order.Status = "Canceled"
		a.db.UpdateStatus(&order)
		return domain.Order{}, paymentErr
	}

	order.Status = "Paid"
	a.db.UpdateStatus(&order)

	deliveryDays, shippingErr := a.shipping.CreateShipping(&order)
	if shippingErr != nil {
		log.Printf("Erro ao criar shipping para o pedido %d: %v", order.ID, shippingErr)
	} else {
		log.Printf("Shipping criado para o pedido %d. Prazo de entrega: %d dias", order.ID, deliveryDays)
	}

	return order, nil
}
