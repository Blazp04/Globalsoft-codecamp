package errors

import "fmt"

type WithdrawError struct {
	Owner   string
	Amount  float64
	Message string
}

func (w *WithdrawError) Error() string {
	return fmt.Sprintf("Cannot withdraw %f from %s. %s", w.Amount, w.Owner, w.Message)
}
