package customErrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	// Generics
	ErrorNotFound             error = errors.New("resource not found")
	ErrorInternalServerError  error = errors.New("internal server error")
	ErrorConflict             error = errors.New("conflict occurred")
	ErrorBadRequest           error = errors.New("bad request")
	ErrorUnprocessableContent error = errors.New("unprocessable content")
)

type ValidationError struct {
	Messages []string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("There were some errors validating:  %s", strings.Join(ve.Messages, ", "))
}
func HandleSqlError(err error) error{
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1452: // BuyerId, carrier id, orderstatusid or warehouseid not found
			return ErrorNotFound
		case 1062: // Duplicated Order Number
			return ErrorConflict
		default:
			return err
		}
	}
	return err
}
