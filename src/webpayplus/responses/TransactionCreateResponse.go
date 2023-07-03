package responses

type TransactionCreateResponse struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}
