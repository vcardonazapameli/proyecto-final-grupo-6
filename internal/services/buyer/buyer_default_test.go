package buyer

import (
	"testing"

	
	mockRepo "github.com/arieleon_meli/proyecto-final-grupo-6/internal/repositories/buyer"
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Buyer_Read(t *testing.T){
	t.Run("find_all", func(t *testing.T) {
		//Arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service :=  NewBuyerDefault(mockRepo)

		expectedBuyers := []models.BuyerDocResponse{
			{
				Id:           1,
			CardNumberId: 123456,
			FirstName:    "John",
			LastName:     "Doe",
		},
		{
			Id:           2,
			CardNumberId: 654321,
			FirstName:    "Jane",
			LastName:     "Smith",
		},
		{
			Id:           3,
			CardNumberId: 789012,
			FirstName:    "Alice",
			LastName:     "Johnson",
		},
		}
		mockRepo.On("GetAll").Return(expectedBuyers, nil)
		//Act
		result, err := service.GetAll()
		//Assert
		assert.NoError(t, err)
		assert.Equal(t, result, expectedBuyers)
		mockRepo.AssertExpectations(t)
	})
	t.Run("find_all_with_error", func(t *testing.T) {
		//Arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service :=  NewBuyerDefault(mockRepo)


		mockRepo.On("GetAll").Return(nil, customErrors.ErrorInternalServerError)
		//Act
		result, err := service.GetAll()
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("find_by_id_existent", func(t *testing.T) {
		//Arrage
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		expectedBuyer := &models.BuyerDocResponse{
			Id: 1,
			CardNumberId: 1234,
			FirstName: "Alberto",
			LastName: "Orozco",
		}
		mockRepo.On("GetById", 1).Return(expectedBuyer, nil)
		//Act
		result, err := service.GetById(1)
		//Assert
		assert.NoError(t,err)
		assert.Equal(t, result, expectedBuyer)
		mockRepo.AssertExpectations(t)

	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Arrage
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		expectedError := customErrors.ErrorNotFound
		mockRepo.On("GetById", 1).Return(nil, customErrors.ErrorNotFound)
		//Act
		result, err := service.GetById(1)
		//Assert
		assert.ErrorAs(t, err, &expectedError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)

	})
	t.Run("get_purchases_reports", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		cardNumberId := 234532
		expectedReports := []models.PurchaseOrderReport{
			{
				BuyerDocResponse: models.BuyerDocResponse {
					Id:           2,
					CardNumberId: 234532,
					FirstName:    "Jacinto",
					LastName:     "Orejuela",
				},
				PurchaseOrdersCount: 2,
			},
		}
		
		mockRepo.On("ValidateCardNumberId",234532).Return(true)
		mockRepo.On("GetPurchasesReports", cardNumberId).Return(expectedReports, nil)
		//Act
		result, err := service.GetPurchasesReports(cardNumberId)
		//Assert
		assert.NoError(t, err)
		assert.Equal(t, result, expectedReports)
		mockRepo.AssertExpectations(t)
	})
	t.Run("get_purchases_reports_non_existent", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		cardNumberId := 234532
		
		mockRepo.On("ValidateCardNumberId",cardNumberId).Return(false)
		
		//Act
		result, err := service.GetPurchasesReports(cardNumberId)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("get_purchases_reports_error_from_repo", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		cardNumberId := 234532
		
		mockRepo.On("ValidateCardNumberId",cardNumberId).Return(true)
		mockRepo.On("GetPurchasesReports",cardNumberId).Return(nil, customErrors.ErrorInternalServerError)
		//Act
		result, err := service.GetPurchasesReports(cardNumberId)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func Test_Buyer_Create(t *testing.T){
	t.Run("create_ok", func(t *testing.T) {
		//arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service:= NewBuyerDefault(mockRepo)
		buyerRequest := models.BuyerDocRequest{
			CardNumberId: 1234,
			FirstName: "Alberto",
			LastName: "Orozco"}
		expectedBuyer := &models.BuyerDocResponse{
			Id: 0,
			CardNumberId: 1234,
			FirstName: "Alberto",
			LastName: "Orozco",
		}
		mockRepo.On("ValidateCardNumberId", 1234).Return(false)
		mockRepo.On("CreateBuyer", mock.Anything).Return(nil)
		//Act
		result, err := service.CreateBuyer(buyerRequest)
		//Assert
		assert.NoError(t,err)
		assert.Equal(t, result, expectedBuyer)
		mockRepo.AssertExpectations(t)
	})
	t.Run("create_conflict", func(t *testing.T) {
		//arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service:= NewBuyerDefault(mockRepo)
		buyerRequest := models.BuyerDocRequest{
			CardNumberId: 1234,
			FirstName: "Alberto",
			LastName: "Orozco"}
		expectedError := customErrors.ErrorConflict
		mockRepo.On("ValidateCardNumberId", 1234).Return(true)
		//Act
		result, err := service.CreateBuyer(buyerRequest)
		//Assert
		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("create_fail", func(t *testing.T) {

		//arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service:= NewBuyerDefault(mockRepo)
		buyerRequest := models.BuyerDocRequest{
			CardNumberId: 0,
			FirstName: "",
			LastName: ""}
		expectedError := customErrors.ErrorUnprocessableContent
		
		//Act
		result, err := service.CreateBuyer(buyerRequest)
		//Assert
		assert.ErrorIs(t, err, expectedError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("create_fail_from_repo", func(t *testing.T) {

		//arrange
		mockRepo:= new(mockRepo.BuyerRepositoryMock)
		service:= NewBuyerDefault(mockRepo)
		buyerRequest := models.BuyerDocRequest{
			CardNumberId: 1234,
			FirstName: "asd",
			LastName: "vxcv",
		}
		mockRepo.On("ValidateCardNumberId", 1234).Return(false)
		mockRepo.On("CreateBuyer", mock.Anything).Return(customErrors.ErrorInternalServerError)
		
		//Act
		result, err := service.CreateBuyer(buyerRequest)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
func Test_Buyer_Update(t *testing.T){
	t.Run("update_existent", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service  := NewBuyerDefault(mockRepo)
		idBuyerToUpdate := 1
		cardNumberId:= 1234
		firstName := "Enrique"
		lastName := "Gomez"
		buyerRequest := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId, 
			FirstName: &firstName,
			LastName:  &lastName,
		}
		buyerToUpdate := &models.BuyerDocResponse{
			Id: idBuyerToUpdate,
			CardNumberId: 4567,
			FirstName: "Carlos",
			LastName: "Bolaños",
		}
		updatedBuyer := &models.BuyerDocResponse{
			Id: idBuyerToUpdate,
			CardNumberId: 1234,
			FirstName: "Enrique",
			LastName: "Gomez",
		}
		buyerDoc := &models.BuyerDocRequest{
			CardNumberId: 1234,
			FirstName: "Enrique",
			LastName: "Gomez",
		}
		mockRepo.On("GetById",idBuyerToUpdate).Return(buyerToUpdate, nil)
		mockRepo.On("ValidateCardNumberIdToUpdate", 1234, idBuyerToUpdate).Return(false)
		mockRepo.On("UpdateBuyer", idBuyerToUpdate, buyerDoc).Return(nil)
		//Act
		result, err := service.UpdateBuyer(idBuyerToUpdate,buyerRequest)
		//Assert
		assert.Equal(t, result, updatedBuyer)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		
	})
	t.Run("update_non_existent", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service  := NewBuyerDefault(mockRepo)
		idBuyerToUpdate := 4
		cardNumberId:= 1234
		firstName := "Enrique"
		lastName := "Gomez"
		buyerRequest := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId, 
			FirstName: &firstName,
			LastName:  &lastName,
		}
		
		mockRepo.On("GetById",idBuyerToUpdate).Return(nil, customErrors.ErrorNotFound)
		
		//Act
		result, err := service.UpdateBuyer(idBuyerToUpdate,buyerRequest)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorNotFound)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("update_conflict", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service  := NewBuyerDefault(mockRepo)
		idBuyerToUpdate := 4
		cardNumberId:= 1234
		firstName := "Enrique"
		lastName := "Gomez"
		buyerRequest := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId, 
			FirstName: &firstName,
			LastName:  &lastName,
		}
		buyerToUpdate := &models.BuyerDocResponse{
			Id: idBuyerToUpdate,
			CardNumberId: 4567,
			FirstName: "Carlos",
			LastName: "Bolaños",
		}
		
		mockRepo.On("GetById",idBuyerToUpdate).Return(buyerToUpdate, nil)
		mockRepo.On("ValidateCardNumberIdToUpdate",1234,4).Return(true)
		
		//Act
		result, err := service.UpdateBuyer(idBuyerToUpdate,buyerRequest)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorConflict)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("update_with_error_from_repo", func(t *testing.T) {
		//Arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service  := NewBuyerDefault(mockRepo)
		idBuyerToUpdate := 1
		cardNumberId:= 1234
		firstName := "Enrique"
		lastName := "Gomez"
		buyerRequest := models.UpdateBuyerDto{
			CardNumberId: &cardNumberId, 
			FirstName: &firstName,
			LastName:  &lastName,
		}
		buyerToUpdate := &models.BuyerDocResponse{
			Id: idBuyerToUpdate,
			CardNumberId: 4567,
			FirstName: "Carlos",
			LastName: "Bolaños",
		}
		buyerDoc := &models.BuyerDocRequest{
			CardNumberId: 1234,
			FirstName: "Enrique",
			LastName: "Gomez",
		}
		mockRepo.On("GetById",idBuyerToUpdate).Return(buyerToUpdate, nil)
		mockRepo.On("ValidateCardNumberIdToUpdate", 1234, idBuyerToUpdate).Return(false)
		mockRepo.On("UpdateBuyer", idBuyerToUpdate, buyerDoc).Return(customErrors.ErrorInternalServerError)
		//Act
		result, err := service.UpdateBuyer(idBuyerToUpdate,buyerRequest)
		//Assert
		assert.ErrorIs(t, err, customErrors.ErrorInternalServerError)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		
	})
}
func Test_Buyer_Delete(t *testing.T){
	t.Run("delete_ok", func(t *testing.T) {
		//arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		mockRepo.On("ValidateIfExistsById", 1).Return(true)
		mockRepo.On("DeleteBuyer", 1).Return( nil)
		//act
		err := service.DeleteBuyer(1)
		//assert
		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "ValidateIfExistsById", 1)
		mockRepo.AssertCalled(t, "DeleteBuyer", 1)
		mockRepo.AssertExpectations(t)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		//arrange
		mockRepo := new(mockRepo.BuyerRepositoryMock)
		service := NewBuyerDefault(mockRepo)
		mockRepo.On("ValidateIfExistsById", 1).Return(false)
		//act
		err := service.DeleteBuyer(1)
		//assert
		assert.ErrorIs(t, err,customErrors.ErrorNotFound)
		mockRepo.AssertCalled(t, "ValidateIfExistsById", 1)
		mockRepo.AssertExpectations(t)
	})
	
}



