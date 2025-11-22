package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Chandra5468/basic-ecom/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ProductNoStock     = errors.New("product has not enough stock")
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer id is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("atleast one item is required to be orderd")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx) // a roll back for fail safe even if something crashed

	qtx := s.repo.WithTx(tx)
	// create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}
	// look for the product if exists
	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ProductNoStock
		}

		// create order items in db. If product exists
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCenters,
		})

		if err != nil {
			return repo.Order{}, err
		}
		currentQuantity := product.Quantity - item.Quantity
		// update the product and reduce the quantity
		qtx.UpdateProductQuantity(ctx, repo.UpdateProductQuantityParams{
			Quantity: currentQuantity,
			ID:       product.ID,
		})
	}
	tx.Commit(ctx)
	return order, nil
}
