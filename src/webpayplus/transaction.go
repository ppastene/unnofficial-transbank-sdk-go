package webpayplus

import (
	"encoding/json"
	"strconv"

	"github.com/ppastene/unnofficial-transbank-sdk-go/src/common"
	"github.com/ppastene/unnofficial-transbank-sdk-go/src/webpayplus/responses"
)

type payload struct {
	BuyOrder  string  `json:"buy_order"`
	SessionId string  `json:"session_id"`
	Amount    float64 `json:"amount"`
	ReturnUrl string  `json:"return_url"`
}

type Transaction struct {
	Options common.Options
}

func newPayload(buyOrder, sessionId string, amount float64, returnUrl string) payload {
	return payload{buyOrder, sessionId, amount, returnUrl}
}

func NewTransaction(options common.Options) Transaction {
	return Transaction{options}
}

func (t Transaction) Create(buyOrder, sessionId string, amount float64, returnUrl string) responses.TransactionCreateResponse {
	payload := newPayload(buyOrder, sessionId, amount, returnUrl)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic("Error al convertir a json")
	}
	var response responses.TransactionCreateResponse
	err = json.Unmarshal(common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/", &jsonData, t.Options), &response)
	if err != nil {
		panic("Error al convertir json a struct")
	}
	return response
}

func (t Transaction) Commit(token string) responses.TransactionStatusResponse {
	var data []byte
	var response responses.TransactionStatusResponse
	err := json.Unmarshal(common.HTTPRequest("PUT", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token, &data, t.Options), &response)
	if err != nil {
		panic("Error al convertir json a struct")
	}
	return response
}

func (t Transaction) Refund(token, amount string) responses.TransactionRefundResponse {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		panic("Error al convertir de string a entero")
	}
	data := map[string]interface{}{
		"amount": intAmount,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic("Error al convertir valor a json")
	}
	var response responses.TransactionRefundResponse
	err = json.Unmarshal(common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token+"/refunds", &jsonData, t.Options), &response)
	if err != nil {
		panic("Error al convertir json a struct")
	}
	return response
}

func (t Transaction) Status(token string) responses.TransactionStatusResponse {
	var data []byte
	var response responses.TransactionStatusResponse
	err := json.Unmarshal(common.HTTPRequest("PUT", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token, &data, t.Options), &response)
	if err != nil {
		panic("Error al convertir json a struct")
	}
	return response
}

func (t Transaction) Capture(token, buyOrder, authorizationCode string, captureAmount float64) responses.TransactionCaptureResponse {
	data := map[string]interface{}{
		"buy_order":          buyOrder,
		"authorization_code": authorizationCode,
		"capture_amount":     captureAmount,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic("Error al convertir valor a json")
	}
	var response responses.TransactionCaptureResponse
	err = json.Unmarshal(common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token+"/capture", &jsonData, t.Options), &response)
	if err != nil {
		panic("Error al convertir json a struct")
	}
	return response
}
