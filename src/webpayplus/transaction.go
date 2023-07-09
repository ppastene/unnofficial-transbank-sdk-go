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

func (t Transaction) Create(buyOrder, sessionId string, amount float64, returnUrl string) (responses.TransactionCreateResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return responses.TransactionCreateResponse{}, err
	}
	request, err := common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/", &jsonData, t.Options)
	if err != nil {
		return responses.TransactionCreateResponse{}, err
	}
	var response responses.TransactionCreateResponse
	err = json.Unmarshal(request, &response)
	if err != nil {
		return responses.TransactionCreateResponse{}, err
	}
	return response, nil
}

func (t Transaction) Commit(token string) (responses.TransactionStatusResponse, error) {
	var data []byte
	request, err := common.HTTPRequest("PUT", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token, &data, t.Options)
	if err != nil {
		return responses.TransactionStatusResponse{}, err
	}
	var response responses.TransactionStatusResponse
	err = json.Unmarshal(request, &response)
	if err != nil {
		return responses.TransactionStatusResponse{}, err
	}
	return response, nil
}

func (t Transaction) Refund(token, amount string) (responses.TransactionRefundResponse, error) {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		return responses.TransactionRefundResponse{}, err
	}
	data := map[string]interface{}{
		"amount": intAmount,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return responses.TransactionRefundResponse{}, err
	}
	request, err := common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token+"/refunds", &jsonData, t.Options)
	if err != nil {
		return responses.TransactionRefundResponse{}, err
	}
	var response responses.TransactionRefundResponse
	err = json.Unmarshal(request, &response)
	if err != nil {
		return responses.TransactionRefundResponse{}, err
	}
	return response, nil
}

func (t Transaction) Status(token string) (responses.TransactionStatusResponse, error) {
	var data []byte
	request, err := common.HTTPRequest("PUT", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token, &data, t.Options)
	if err != nil {
		return responses.TransactionStatusResponse{}, err
	}
	var response responses.TransactionStatusResponse
	err = json.Unmarshal(request, &response)
	if err != nil {
		return responses.TransactionStatusResponse{}, err
	}
	return response, nil
}

func (t Transaction) Capture(token, buyOrder, authorizationCode string, captureAmount float64) (responses.TransactionCaptureResponse, error) {
	data := map[string]interface{}{
		"buy_order":          buyOrder,
		"authorization_code": authorizationCode,
		"capture_amount":     captureAmount,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return responses.TransactionCaptureResponse{}, err
	}
	request, err := common.HTTPRequest("POST", "/rswebpaytransaction/api/webpay/v1.3/transactions/"+token+"/capture", &jsonData, t.Options)
	var response responses.TransactionCaptureResponse
	err = json.Unmarshal(request, &response)
	if err != nil {
		return responses.TransactionCaptureResponse{}, err
	}
	return response, nil
}
