package product_test

import (
	"context"
	"errors"
	"github.com/khasmag06/zdravservice-test/internal/service/product"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/khasmag06/zdravservice-test/internal/models"
)

func TestService_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepository(ctrl)
	svc := product.New(mockRepo)

	tests := []struct {
		name          string
		page          int
		limit         int
		sortOrder     string
		repoResult    []models.Product
		repoError     error
		expected      []models.Product
		expectedError error
	}{
		{
			name:      "valid result",
			page:      1,
			limit:     2,
			sortOrder: "asc",
			repoResult: []models.Product{
				{ID: 1, Properties: map[string]string{"штрихкод": "12345678"}},
				{ID: 2, Properties: map[string]string{"штрихкод": "87654321"}},
			},
			repoError: nil,
			expected: []models.Product{
				{ID: 1, Properties: map[string]string{"штрихкод": "12345678"}},
				{ID: 2, Properties: map[string]string{"штрихкод": "87654321"}},
			},
			expectedError: nil,
		},
		{
			name:          "empty result",
			page:          1,
			limit:         10,
			sortOrder:     "asc",
			repoResult:    nil,
			repoError:     nil,
			expected:      []models.Product{},
			expectedError: nil,
		},
		{
			name:          "repo error",
			page:          1,
			limit:         10,
			sortOrder:     "asc",
			repoResult:    nil,
			repoError:     errors.New("repository error"),
			expected:      nil,
			expectedError: errors.New("repository error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.EXPECT().GetProducts(gomock.Any(), test.page, test.limit, test.sortOrder).Return(test.repoResult, test.repoError)

			products, err := svc.GetProducts(context.Background(), test.page, test.limit, test.sortOrder)

			assert.Equal(t, test.expected, products, "Test case %s failed: Products not as expected", test.name)
			assert.Equal(t, test.expectedError, err, "Test case %s failed: Error not as expected", test.name)
		})
	}
}
