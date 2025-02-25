package seller_test

import (
	"testing"

	repo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestReadSellers(t *testing.T) {
	t.Run("Get All Sellers", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetAll").Return(map[int]models.Seller{
			1: {Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			2: {Id: 2, Cid: 2000, CompanyName: "Company2", Address: "San Martin 2", Telephone: "11333444", LocalityID: 2},
		}, nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		expectedDocs := map[int]models.SellerDoc{
			1: {Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			2: {Id: 2, Cid: 2000, CompanyName: "Company2", Address: "San Martin 2", Telephone: "11333444", LocalityID: 2},
		}

		// Act
		sDocs, err := sv.GetAll()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, sDocs, 2)
		assert.Equal(t, expectedDocs, sDocs)
	})

	t.Run("Get Seller By ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetByID", 1).Return(models.Seller{
			Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		expectedDoc := models.SellerDoc{
			Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1}

		// Act
		sDoc, err := sv.GetByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedDoc, sDoc)
	})
}
