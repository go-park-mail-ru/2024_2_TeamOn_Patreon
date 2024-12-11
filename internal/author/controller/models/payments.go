package models

// Модель выплат
type Payments struct {
	// Сумма выплат
	Amount int `json:"amount"`
}

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

type PaymentResponse struct {
	ID           string           `json:"id"`
	Status       string           `json:"status"`
	Paid         bool             `json:"paid"`
	Amount       Amount           `json:"amount"`
	CreatedAt    string           `json:"created_at"`
	Description  string           `json:"description"`
	Confirmation ConfirmationResp `json:"confirmation"`
}

type ConfirmationResp struct {
	Type            string `json:"type"`
	ConfirmationURL string `json:"confirmation_url"`
}
