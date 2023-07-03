package responses

type TransactionStatusResponse struct {
	Vci                string      `json:"vci"`
	Amount             float64     `json:"amount"`
	Status             string      `json:"status"`
	BuyOrder           string      `json:"buy_order"`
	SessionId          string      `json:"session_id"`
	CardDetail         cardDetails `json:"card_detail"`
	AccountingDate     string      `json:"accounting_date"`
	TransactionDate    string      `json:"transaction_date"`
	AuthorizationCode  string      `json:"authorization_code"`
	PaymentTypeCode    string      `json:"payment_type_code"`
	ResponseCode       int         `json:"response_code"`
	InstallmentsAmount int         `json:"installments_amount"`
	InstallmentsNumber int         `json:"installments_number"`
	Balance            int         `json:"balance"`
}

type cardDetails struct {
	CardNumber string `json:"card_number"`
}

func (t TransactionStatusResponse) IsApproved() bool {
	if t.ResponseCode != 0 {
		return false
	}
	switch t.Status {
	case "PARTIALLY_NULLIFIED", "REVERSED", "NULLIFIED", "AUTHORIZED":
		return true
	default:
		return false
	}
}
