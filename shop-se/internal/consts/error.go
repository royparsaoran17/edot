package consts

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrProductAlreadyExist = Error("product already exist")
	ErrShopAlreadyExist    = Error("shop already exist")
	ErrShopStockEmpty      = Error("shop stock is not empty")
	ErrShopNotFound        = Error("shop not found")
	ErrUserNotFound        = Error("user not found")

	ErrWrongPassword     = Error("wrong password")
	ErrPhoneAlreadyExist = Error("phone already exist")
	ErrProductNotFound   = Error("product not found")
	ErrStockNotFound     = Error("stock not found")

	ErrInvalidUUID = Error("UUID is not in its proper form")
)
