//go:generate mockgen -source=$GOFILE -destination=mocks_test.go -package=$GOPACKAGE
package api

import (
	"context"
	"github.com/khasmag06/zdravservice-test/internal/models"
)

type productsService interface {
	GetProducts(ctx context.Context, page int, limit int, sortOrder string) ([]models.Product, error)
}

type logger interface {
	Info(text ...any)
	Error(text ...any)
	Errorf(format string, args ...any)
}
