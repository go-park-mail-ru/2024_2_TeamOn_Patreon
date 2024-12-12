package models

type InfoPaySubscription struct {
	AuthorID    string // Для редиректа
	Cost        string // Счёт на оплату
	Description string // Описание платежа
}

type PaymentRequest struct {
	Amount       Amount          `json:"amount"`
	Description  string          `json:"description"`
	Confirmation ConfirmationReq `json:"confirmation"`
	Metadata     interface{}     `json:"metadata,omitempty"`
	Test         bool            `json:"test"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type ConfirmationReq struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url,omitempty"`
}

type NotificationPaymentService struct {
	Type   string          `json:"type"`
	Event  string          `json:"event"`
	Object PaymentResponse `json:"object"`
}

type PaymentResponse struct {
	ID           string           `json:"id"`
	Status       string           `json:"status"`
	Paid         bool             `json:"paid"`
	Amount       Amount           `json:"amount"`
	CreatedAt    string           `json:"created_at"`
	Description  string           `json:"description"`
	Confirmation ConfirmationResp `json:"confirmation"`
	// Metadata     string           `json:"metadata"`
}

type ConfirmationResp struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
}
