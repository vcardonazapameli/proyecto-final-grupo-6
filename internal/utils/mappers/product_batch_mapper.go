package mappers

import (
	"time"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/customErrors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func parseDateTime(dateStr, hourStr string) (time.Time, error) {
	if hourStr == "" {
		hourStr = "00"
	}

	if len(hourStr) <= 2 {
		hourStr = hourStr + ":00:00"
	}

	dateTimeStr := dateStr + " " + hourStr
	return time.Parse("2006-01-02 15:04:05", dateTimeStr)
}

func ProductBatchRequestToProductBatch(request models.ProductBatchRequest) (*models.ProductBatchResponse, error) {
	dueDate, err := parseDateTime(request.DueDate, "")
	if err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}

	manufacturingDate, err := parseDateTime(request.ManufacturingDate, "")
	if err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}

	manufacturingHour, err := parseDateTime(request.ManufacturingDate, request.ManufacturingHour)
	if err != nil {
		return nil, customErrors.ErrorUnprocessableContent
	}

	return &models.ProductBatchResponse{
		BatchNumber:        request.BatchNumber,
		CurrentQuantity:    request.CurrentQuantity,
		CurrentTemperature: request.CurrentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    request.InitialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: request.MinimumTemperature,
		ProductId:          request.ProductId,
		SectionId:          request.SectionId,
	}, nil
}
