package product

import (
	"testing"

	productRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product"
	productTypeRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/product_type"
	sellerRepository "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/seller"
	customErrors "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	errorCustom "github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Product_Create(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)

		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		expectedProduct := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("ExistInDb", productRequest.ProductCode).Return(false, nil)
		productTypeRepository.On("GetById", productRequest.ProductType).Return(&models.ProductTypeDocResponse{}, nil)
		sellerRepository.On("GetByID", productRequest.Seller).Return(models.Seller{}, nil)
		productRepository.On("Create", mock.Anything).Return(nil, nil)

		// Act
		existInDb, _ := productRepository.ExistInDb(productRequest.ProductCode)
		product, err := productService.Create(productRequest)
		product.Id = 1

		// Assert
		require.False(t, existInDb)
		require.NoError(t, err)
		require.Equal(t, expectedProduct, product)
		productRepository.AssertExpectations(t)
		productTypeRepository.AssertExpectations(t)
		sellerRepository.AssertExpectations(t)
	})
	t.Run("create_conflict", func(t *testing.T) {
		// Arrange
		expectedError := customErrors.ErrorConflict
		productRepository := new(productRepository.ProductRepositoryMock)
		productRequest := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("Create", productRequest).Return(expectedError)

		// Act
		err := productRepository.Create(productRequest)

		// Assert
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
	})
	t.Run("create_validation_error", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)

		productRequest := models.ProductDocRequest{
			ProductCode:                    "",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		// Act
		product, err := productService.Create(productRequest)

		// Assert
		require.Error(t, err)
		require.Nil(t, product)
	})

	t.Run("create_conflict_product_exists", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)

		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("ExistInDb", productRequest.ProductCode).Return(true, nil)

		// Act
		product, err := productService.Create(productRequest)

		// Assert
		require.Nil(t, product)
		require.Error(t, err)
		require.Equal(t, errorCustom.ErrorConflict, err)
		productRepository.AssertExpectations(t)
	})

	t.Run("create_invalid_product_type", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)

		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("ExistInDb", productRequest.ProductCode).Return(false, nil)
		productTypeRepository.On("GetById", productRequest.ProductType).Return(nil, nil)

		// Act
		product, err := productService.Create(productRequest)

		// Assert
		require.Nil(t, product)
		require.Error(t, err)
		require.Equal(t, errorCustom.ErrorConflict, err)
		productRepository.AssertExpectations(t)
		productTypeRepository.AssertExpectations(t)
	})

	t.Run("create_invalid_seller", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)
		expectedError := customErrors.ErrorConflict

		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("ExistInDb", productRequest.ProductCode).Return(false, nil)
		productTypeRepository.On("GetById", productRequest.ProductType).Return(&models.ProductTypeDocResponse{}, nil)
		sellerRepository.On("GetByID", productRequest.Seller).Return(models.Seller{}, expectedError)

		// Act
		product, err := productService.Create(productRequest)

		// Assert
		require.Nil(t, product)
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
		productTypeRepository.AssertExpectations(t)
		sellerRepository.AssertExpectations(t)
	})
	t.Run("create_repository_error", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		sellerRepository := new(sellerRepository.SellerRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, sellerRepository)
		expectedError := customErrors.ErrorInternalServerError

		productRequest := models.ProductDocRequest{
			ProductCode:                    "PROD004",
			Description:                    "Producto D",
			ExpirationRate:                 0.7,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          2,
			Height:                         5,
			Length:                         10,
			NetWeight:                      0.5,
			ProductType:                    2,
			Seller:                         2,
		}

		productRepository.On("ExistInDb", productRequest.ProductCode).Return(false, nil)
		productTypeRepository.On("GetById", productRequest.ProductType).Return(&models.ProductTypeDocResponse{}, nil)
		sellerRepository.On("GetByID", productRequest.Seller).Return(models.Seller{}, nil)
		productRepository.On("Create", mock.Anything).Return(expectedError)

		// Act
		product, err := productService.Create(productRequest)

		// Assert
		require.Nil(t, product)
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
		productTypeRepository.AssertExpectations(t)
		sellerRepository.AssertExpectations(t)
	})
}

func Test_Product_Read(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedProducts := []models.ProductDocResponse{
			{
				Id:                             1,
				ProductCode:                    "P001",
				Description:                    "Producto A",
				ExpirationRate:                 365,
				RecommendedFreezingTemperature: -18,
				FreezingRate:                   0.5,
				Width:                          10,
				Height:                         20,
				Length:                         30,
				NetWeight:                      1,
				ProductType:                    1,
				Seller:                         1,
			},
			{
				Id:                             2,
				ProductCode:                    "P002",
				Description:                    "Producto B",
				ExpirationRate:                 365,
				RecommendedFreezingTemperature: -18,
				FreezingRate:                   0.5,
				Width:                          10,
				Height:                         20,
				Length:                         30,
				NetWeight:                      1,
				ProductType:                    2,
				Seller:                         2,
			},
		}
		productRepository.On("GetAll").Return(expectedProducts, nil)

		// Act
		products, err := productService.GetAll()

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedProducts, products)
		productRepository.AssertCalled(t, "GetAll")
		productRepository.AssertExpectations(t)
	})
	t.Run("find_all_fail", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorInternalServerError
		productRepository.On("GetAll").Return(nil, expectedError)

		// Act
		product, err := productService.GetAll()

		// Assert
		require.Error(t, err)
		require.Nil(t, product)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertCalled(t, "GetAll")
		productRepository.AssertExpectations(t)
	})
	t.Run("find_by_id_existent", func(t *testing.T) {
		//Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)

		expectedProduct := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001",
			Description:                    "Producto A",
			ExpirationRate:                 365,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   0.5,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		productRepository.On("GetById", 1).Return(expectedProduct, nil)

		//Act
		product, err := productService.GetById(1)

		//Assert
		require.NoError(t, err)
		require.Equal(t, expectedProduct, product)
		productRepository.AssertExpectations(t)
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorNotFound
		productRepository.On("GetById", 2).Return(nil, expectedError)

		//Act
		product, err := productService.GetById(2)

		//Assert
		require.Error(t, err)
		require.Nil(t, product)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
	})
}

func Test_Product_Update(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productRequest := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001B",
			Description:                    "Producto B",
			ExpirationRate:                 200,
			RecommendedFreezingTemperature: -20,
			FreezingRate:                   -20,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		productRepository.On("Update", 1, productRequest).Return(nil)

		// Act
		err := productRepository.Update(1, productRequest)

		// Assert
		require.NoError(t, err)
		productRepository.AssertExpectations(t)
	})
	t.Run("update_non_existent", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorNotFound

		productRequest := models.ProductUpdateDocRequest{
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		productRepository.On("GetById", 1).Return(nil, expectedError)

		// Act
		product, err := productService.Update(1, productRequest)

		// Assert
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		require.Nil(t, product)
		productRepository.AssertExpectations(t)
	})
	t.Run("update_conflict", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)

		productRequest := models.ProductUpdateDocRequest{
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		existingProduct := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001",
			Description:                    "Producto A",
			ExpirationRate:                 100,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   -0.5,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		productRepository.On("GetById", 1).Return(existingProduct, nil)
		productRepository.On("MatchWithTheSameProductCode", 1, *productRequest.ProductCode).Return(true, nil)

		// Act
		_, err := productService.Update(1, productRequest)

		// Assert
		require.Error(t, err)
		require.Equal(t, customErrors.ErrorConflict, err)
		productRepository.AssertExpectations(t)
	})
	t.Run("update_product_code_conflict", func(t *testing.T) {
		// Arrange}
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)

		productRequest := models.ProductUpdateDocRequest{
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		existingProduct := &models.ProductDocResponse{
			Id:                             1,
			ProductCode:                    "P001",
			Description:                    "Producto A",
			ExpirationRate:                 100,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   -0.5,
			Width:                          10,
			Height:                         20,
			Length:                         30,
			NetWeight:                      1,
			ProductType:                    1,
			Seller:                         1,
		}

		productRepository.On("GetById", 1).Return(existingProduct, nil)
		productRepository.On("MatchWithTheSameProductCode", 1, *productRequest.ProductCode).Return(true, nil)

		// Act
		_, err := productService.Update(1, productRequest)

		// Assert
		require.Error(t, err)
		require.Equal(t, customErrors.ErrorConflict, err)
		productRepository.AssertExpectations(t)
	})
	t.Run("update_product_validation_fails", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		productService := NewProductService(productRepository, productTypeRepository, nil)
		productRequest := models.ProductUpdateDocRequest{
			ProductType:    new(int),
			ProductCode:    new(string),
			Description:    new(string),
			ExpirationRate: new(float64),
		}
		*productRequest.ProductType = 2
		*productRequest.ProductCode = "P001B"
		*productRequest.Description = "Producto B"
		*productRequest.ExpirationRate = 200

		productTypeRepository.On("GetById", *productRequest.ProductType).Return(&models.ProductTypeDocResponse{}, nil)
		productRepository.On("GetById", 1).Return(&models.ProductDocResponse{Id: 1}, nil)
		productRepository.On("MatchWithTheSameProductCode", 1, *productRequest.ProductCode).Return(false, nil)

		// Act
		product, err := productService.Update(1, productRequest)

		// Assert
		require.Error(t, err)
		require.Nil(t, product)
		productTypeRepository.AssertExpectations(t)
		productRepository.AssertExpectations(t)
	})
	t.Run("update_product_type_not_found", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productTypeRepository := new(productTypeRepository.ProductTypeRepositoryMock)
		expectedError := customErrors.ErrorNotFound

		productRequest := models.ProductUpdateDocRequest{
			ProductType: new(int),
		}
		*productRequest.ProductType = 2

		productRepository.On("GetById", 1).Return(&models.ProductDocResponse{}, nil)
		productTypeRepository.On("GetById", *productRequest.ProductType).Return(nil, expectedError)

		// Act
		product, _ := productRepository.GetById(1)
		productType, err := productTypeRepository.GetById(*productRequest.ProductType)

		// Assert
		require.NotNil(t, product)
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		require.Nil(t, productType)
		productTypeRepository.AssertExpectations(t)
		productRepository.AssertExpectations(t)
	})
}

func Test_Product_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		//Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)

		productRepository.On("GetById", 1).Return(&models.ProductDocResponse{Id: 1}, nil)
		productRepository.On("Delete", 1).Return(nil)

		//Act
		err := productService.Delete(1)

		//Assert
		require.NoError(t, err)
		productRepository.AssertExpectations(t)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		//Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorNotFound

		productRepository.On("GetById", 2).Return(nil, expectedError)

		//Act
		err := productService.Delete(2)

		//Assert
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertCalled(t, "GetById", 2)
		productRepository.AssertExpectations(t)
	})
	t.Run("delete_repository_error", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorInternalServerError

		productRepository.On("GetById", 1).Return(&models.ProductDocResponse{Id: 1}, nil)
		productRepository.On("Delete", 1).Return(expectedError)

		// Act
		err := productService.Delete(1)

		// Assert
		require.Error(t, err)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
	})
}
func Test_Product_GetProductRecords(t *testing.T) {
	t.Run("get_records_ok", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)

		expectedRecords := []models.ProductRecordByProductResponse{
			{
				ProductId:    1,
				Description:  "Producto A",
				RecordsCount: 1,
			},
			{
				ProductId:    2,
				Description:  "Producto B",
				RecordsCount: 2,
			},
			{
				ProductId:    3,
				Description:  "Producto C",
				RecordsCount: 3,
			},
		}

		productRepository.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return(expectedRecords, nil)

		// Act
		records, err := productService.GetProductRecords(nil, nil, nil)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedRecords, records)
		productRepository.AssertExpectations(t)
	})

	t.Run("get_records_error", func(t *testing.T) {
		// Arrange
		productRepository := new(productRepository.ProductRepositoryMock)
		productService := NewProductService(productRepository, nil, nil)
		expectedError := customErrors.ErrorInternalServerError

		productRepository.On("GetProductRecords", mock.Anything, mock.Anything, mock.Anything).Return(nil, expectedError)

		// Act
		records, err := productService.GetProductRecords(nil, nil, nil)

		// Assert
		require.Error(t, err)
		require.Nil(t, records)
		require.ErrorAs(t, err, &expectedError)
		productRepository.AssertExpectations(t)
	})
}
