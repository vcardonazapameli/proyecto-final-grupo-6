package validators

import (
	"errors"
	"reflect"
	"strings"
)

func ValidateNoEmptyFields(v any) error {
	val := reflect.ValueOf(v)
	var invalidFields []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		if field.IsZero() {
			invalidFields = append(invalidFields, fieldType.Tag.Get("json"))
		}
	}
	if len(invalidFields) > 0 {
		return errors.New(strings.Join(invalidFields, ", ") + " cannot be empty")
	}
	return nil
}

func UpdateEntity[T any](dto any, entity T) T {
	entityValue := reflect.ValueOf(entity).Elem()
	dtoValue := reflect.ValueOf(dto)
	for i := 0; i < dtoValue.NumField(); i++ {
		dtoField := dtoValue.Field(i)
		fieldType := dtoValue.Type().Field(i)
		entityField := entityValue.FieldByName(fieldType.Name)
		if entityField.IsValid() && entityField.CanSet() {
			if dtoField.Kind() == entityField.Kind() {
				entityField.Set(dtoField)
			}
		}
	}
	return entity
}
