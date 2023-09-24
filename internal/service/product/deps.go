//go:generate mockgen -source=$GOFILE -destination=mocks_test.go -package=$GOPACKAGE
package product

import (
	"context"
	"github.com/khasmag06/zdravservice-test/internal/models"
)

type repository interface {
	GetProducts(ctx context.Context, page int, limit int, sortOrder string) ([]models.Product, error)
}
