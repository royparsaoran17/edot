package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type PaymentCreate struct {
	ID      string    `json:"id"`
	OrderID string    `json:"order_id"`
	Method  string    `json:"method"`
	Status  string    `json:"status"`
	Amount  float64   `json:"amount"`
	PaidAt  time.Time `json:"paid_at"`
}

func (r *PaymentCreate) Validate() error {
	return validation.Errors{
		"order_id": validation.Validate(&r.OrderID, validation.Required),
		"method":   validation.Validate(&r.Method, validation.Required),
		"status":   validation.Validate(&r.Status, validation.Required),
		"amount":   validation.Validate(&r.Amount, validation.Required),
		"paid_at":  validation.Validate(&r.PaidAt, validation.Required),
	}.Filter()
}

type OrderPayment struct {
	OrderID       string  `json:"-"` // dari path param
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
}

func (p *OrderPayment) Validate() error {
	return validation.Errors{
		"payment_method": validation.Validate(p.PaymentMethod, validation.Required, validation.In("VA_BCA", "VA_BNI", "VA_BRI", "EWALLET", "CREDIT_CARD")),
		"amount":         validation.Validate(p.Amount, validation.Required, validation.Min(1.0)),
	}.Filter()
}
