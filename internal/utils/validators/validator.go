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

func UpdateEntityFromDTO[T any](dto any, entity T) T {
	entityValue := reflect.ValueOf(entity).Elem()
	entityType := entityValue.Type()
	newEntityPtr := reflect.New(entityType)
	newEntity := newEntityPtr.Elem()
	newEntity.Set(entityValue)
	dtoValue := reflect.ValueOf(dto)
	if dtoValue.NumField() == 0 {
		return newEntityPtr.Interface().(T)
	}
	for i := 0; i < dtoValue.NumField(); i++ {
		dtoField := dtoValue.Field(i)
		fieldType := dtoValue.Type().Field(i)
		if dtoField.Kind() == reflect.Ptr && !dtoField.IsNil() {
			entityField := newEntity.FieldByName(fieldType.Name)
			if entityField.CanSet() {
				entityField.Set(dtoField.Elem())
			}
		}
	}
	return newEntityPtr.Interface().(T)
}
