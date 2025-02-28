package seller_test

import (
	"testing"

	repo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/services/seller"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestServiceSeller_Read(t *testing.T) {
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

	t.Run("Find by non existent ID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetByID", 9999).Return(models.Seller{}, customErrors.ErrorNotFound)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		res, err := sv.GetByID(9999)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)

	})

	t.Run("Get All Fails: Internal Error", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetAll").Return(map[int]models.Seller{}, customErrors.ErrorInternalServerError)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		sDocs, err := sv.GetAll()

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Len(t, sDocs, 0)
	})
}

func TestServiceSeller_Create(t *testing.T) {
	t.Run("Create Seller Correctly", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		newSeller := models.NewSeller(1, 1000, "Company1", "Avenida Falsa 1", "11222333", 1)

		mockRepo.On("Save", newSeller).Return(nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		expected := models.NewSellerDoc(1, 1000, "Company1", "Avenida Falsa 1", "11222333", 1)

		// Act
		res, err := sv.Create(*models.NewSellerDoc(1, 1000, "Company1", "Avenida Falsa 1", "11222333", 1))

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, *expected, res)
	})

	t.Run("Create Seller With Invalid attributes", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		res, err := sv.Create(*models.NewSellerDoc(1, 0, "", "", "", 1))

		// Assert
		var valerr customErrors.ValidationError
		assert.ErrorAs(t, err, &valerr)
		assert.Empty(t, res)
	})

	t.Run("Create Seller With Invalid Locality ID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("Save", models.NewSeller(1, 1000, "Company1", "Avenida Falsa 1", "11222333", -1)).Return(customErrors.ErrorNotFound)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		res, err := sv.Create(*models.NewSellerDoc(1, 1000, "Company1", "Avenida Falsa 1", "11222333", -1))

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)
	})

	t.Run("Create Seller With Duplicated CID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("Save", models.NewSeller(1, 1000, "Company1", "Avenida Falsa 1", "11222333", 1)).Return(customErrors.ErrorConflict)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		res, err := sv.Create(*models.NewSellerDoc(1, 1000, "Company1", "Avenida Falsa 1", "11222333", 1))

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorConflict)
		assert.Empty(t, res)
	})

}

func TestServiceSeller_Update(t *testing.T) {
	t.Run("Update Seller Successfully: no-cid given", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)

		mockRepo.On("Update", *models.NewSeller(1, 2000, "New Name", "New Address", "11333444", 2)).Return(nil)
		mockRepo.On("SearchByCID", 2000).Return(models.Seller{}, false)
		mockRepo.On("GetByID", 1).Return(models.Seller{
			Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		var (
			newCID        int    = 2000
			newName       string = "New Name"
			newAddr       string = "New Address"
			newPhone      string = "11333444"
			newLocalityId int    = 2
		)

		expectedSDoc := models.NewSellerDoc(1, 2000, "New Name", "New Address", "11333444", 2)

		// Act
		res, err := sv.Update(1, &newCID, &newName, &newAddr, &newPhone, &newLocalityId)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, *expectedSDoc, res)

	})

	t.Run("Update Seller Fails: trying to update with an existing CID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)

		mockRepo.On("Update", *models.NewSeller(1, 1000, "New Name", "New Address", "11222333", 1)).Return(nil)

		mockRepo.On("GetByID", 1).Return(models.Seller{
			Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			nil)

		mockRepo.On("SearchByCID", 2000).Return(models.Seller{
			Id: 2, Cid: 2000, CompanyName: "Company2", Address: "San Martin 2", Telephone: "11222333", LocalityID: 1},
			true)

		sv := seller.NewSellerServiceDefault(mockRepo)

		var cid int = 2000

		// Act
		res, err := sv.Update(1, &cid, nil, nil, nil, nil)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorConflict)
		assert.Empty(t, res)

	})

	t.Run("Update Seller Fails: NotFound ID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetByID", 9999).Return(models.Seller{}, customErrors.ErrorNotFound)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		res, err := sv.Update(9999, nil, nil, nil, nil, nil)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Empty(t, res)

	})

	t.Run("Update Seller Fails: Validation Error", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("GetByID", 1).Return(models.Seller{
			Id: 1, Cid: 1000, CompanyName: "Company1", Address: "San Martin 1", Telephone: "11222333", LocalityID: 1},
			nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		var wrongFormatCID int = -1
		// Act
		res, err := sv.Update(1, &wrongFormatCID, nil, nil, nil, nil)

		// Assert
		assert.ErrorAs(t, err, &customErrors.ValidationError{})
		assert.Empty(t, res)

	})
}

func TestServiceSeller_Delete(t *testing.T) {
	t.Run("Delete Ok", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("Delete", 1).Return(nil)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		err := sv.Delete(1)

		// Assert
		assert.NoError(t, err)

	})

	t.Run("Delete Failed: Not Found ID", func(t *testing.T) {

		// Arrange
		mockRepo := new(repo.SellerRepositoryMock)
		mockRepo.On("Delete", 9999).Return(customErrors.ErrorNotFound)

		sv := seller.NewSellerServiceDefault(mockRepo)

		// Act
		err := sv.Delete(9999)

		// Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
	})
}
