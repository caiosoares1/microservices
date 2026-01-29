package db

import (
	"context"
	"fmt"

	"github.com/caiosoares1/microservices/shipping/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID      int64
	Status       string
	DeliveryDays int32
}

type ShippingItem struct {
	gorm.Model
	ShippingID  uint
	ProductCode string
	Quantity    int32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Shipping, error) {
	var shippingEntity Shipping
	res := a.db.WithContext(ctx).First(&shippingEntity, id)
	if res.Error != nil {
		return domain.Shipping{}, res.Error
	}

	var itemsEntity []ShippingItem
	a.db.WithContext(ctx).Where("shipping_id = ?", shippingEntity.ID).Find(&itemsEntity)

	items := make([]domain.ShippingItem, len(itemsEntity))
	for i, item := range itemsEntity {
		items[i] = domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		}
	}

	shipping := domain.Shipping{
		ID:           int64(shippingEntity.ID),
		OrderID:      shippingEntity.OrderID,
		Status:       shippingEntity.Status,
		DeliveryDays: shippingEntity.DeliveryDays,
		Items:        items,
		CreatedAt:    shippingEntity.CreatedAt.UnixNano(),
	}
	return shipping, nil
}

func (a Adapter) Save(ctx context.Context, shipping *domain.Shipping) error {
	shippingModel := Shipping{
		OrderID:      shipping.OrderID,
		Status:       shipping.Status,
		DeliveryDays: shipping.DeliveryDays,
	}

	res := a.db.WithContext(ctx).Create(&shippingModel)
	if res.Error != nil {
		return res.Error
	}

	shipping.ID = int64(shippingModel.ID)

	// Salva os itens
	for _, item := range shipping.Items {
		itemModel := ShippingItem{
			ShippingID:  shippingModel.ID,
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		}
		if err := a.db.WithContext(ctx).Create(&itemModel).Error; err != nil {
			return err
		}
	}

	return nil
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("shipping"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Shipping{}, &ShippingItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
