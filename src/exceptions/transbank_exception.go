package exceptions

import "fmt"

type TransbankExpcetion struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   int    `json:"status_code"`
}

func (e TransbankExpcetion) Error() string {
	return fmt.Sprintf("Error message: %s. Status code: %v", e.ErrorMessage, e.StatusCode)
}
