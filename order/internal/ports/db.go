package ports

import domain "github.com/huukhaitran/microservices/order/internal/application/core"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
