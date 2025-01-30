package seller

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	messages []string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("There were some errors validating seller: \n %s", strings.Join(ve.messages, ", "))
}
