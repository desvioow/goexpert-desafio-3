package database

import (
	"database/sql"
	"github.com/desvioow/goexpert-desafio-3/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) ListAll() ([]entity.Order, error) {
	stmt, err := r.Db.Prepare("SELECT * FROM orders")
	if err != nil {
		return []entity.Order{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return []entity.Order{}, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
