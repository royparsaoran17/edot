package consts

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrProductAlreadyExist   = Error("product already exist")
	ErrWarehouseAlreadyExist = Error("warehouse already exist")
	ErrWarehouseNotFound     = Error("warehouse not found")

	ErrWrongPassword     = Error("wrong password")
	ErrPhoneAlreadyExist = Error("phone already exist")
	ErrProductNotFound   = Error("product not found")

	ErrInvalidUUID = Error("UUID is not in its proper form")
)
