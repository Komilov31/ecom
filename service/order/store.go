package cart

import (
	"database/sql"

	"github.com/Komilov31/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES ($1, $2, $3, $4)",
		order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec(`INSET INTO order_items (orderId, productId, quantity, price) 
						VALUES($1, $2, $3, $4)`, orderItem.OrderId, orderItem.ProductId, orderItem.Quantity, orderItem.Price)
	return err
}
