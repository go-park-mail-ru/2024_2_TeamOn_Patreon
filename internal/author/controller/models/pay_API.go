package models

const (
	TypeSubscription = "subscription"
	TypeTip          = "tip"
)

type InfoPaySubscription struct {
	AuthorID    string // Для редиректа
	Cost        string // Счёт на оплату
	Description string // Описание платежа
	PayType     string // Тип платежа
}

//go:generate easyjson -all

//easyjson:json
type PaymentRequest struct {
	Amount       Amount          `json:"amount"`
	Description  string          `json:"description"`
	Confirmation ConfirmationReq `json:"confirmation"`
	Metadata     Metadata        `json:"metadata,omitempty"`
	Test         bool            `json:"test"`
}

//easyjson:json
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

//easyjson:json
type ConfirmationReq struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url,omitempty"`
}

//easyjson:json
type NotificationPaymentService struct {
	Type   string          `json:"type"`
	Event  string          `json:"event"`
	Object PaymentResponse `json:"object"`
}

//easyjson:json
type PaymentResponse struct {
	ID           string           `json:"id"`
	Status       string           `json:"status"`
	Paid         bool             `json:"paid"`
	Amount       Amount           `json:"amount"`
	CreatedAt    string           `json:"created_at"`
	Description  string           `json:"description"`
	Confirmation ConfirmationResp `json:"confirmation"`
	Metadata     Metadata         `json:"metadata"`
}

//easyjson:json
type Metadata struct {
	PayType string `json:"pay_type"` // Поле для типа платежа
}

//easyjson:json
type ConfirmationResp struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
}
