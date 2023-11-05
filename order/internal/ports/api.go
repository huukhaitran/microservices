package ports

import domain "github.com/huukhaitran/microservices/order/internal/application/core"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
