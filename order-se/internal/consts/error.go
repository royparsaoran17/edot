package consts

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrProductAlreadyExist = Error("product already exist")
	ErrOrderAlreadyExist   = Error("order already exist")
	ErrOrderStockEmpty     = Error("order stock is not empty")
	ErrOrderNotFound       = Error("order not found")
	ErrUserNotFound        = Error("user not found")
	ErrDataNotFound        = Error("data not found")

	ErrWrongPassword     = Error("wrong password")
	ErrPhoneAlreadyExist = Error("phone already exist")
	ErrDataAlreadyExist  = Error("data already exist")
	ErrProductNotFound   = Error("product not found")
	ErrProductStockEmpty = Error("product stock empty")
	ErrStockNotFound     = Error("stock not found")

	ErrInvalidUUID = Error("UUID is not in its proper form")
)
