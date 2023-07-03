package common

const ENVIRONMENT_PRODUCTION string = "LIVE"
const ENVIRONMENT_INTEGRATION string = "TEST"

const BASE_URL_PRODUCTION string = "https://webpay3g.transbank.cl"
const BASE_URL_INTEGRATION string = "https://webpay3gint.transbank.cl"

type Options struct {
	CommerceCode string
	ApiKey       string
	Environment  string
}

func newOptions(commerceCode, apiKey string, environment ...string) Options {
	if environment[0] == "" || len(environment[0]) == 0 {
		return Options{commerceCode, apiKey, ENVIRONMENT_INTEGRATION}
	}
	return Options{commerceCode, apiKey, environment[0]}
}

func (o Options) ForProduction(commerceCode, apiKey string) Options {
	return newOptions(commerceCode, apiKey, ENVIRONMENT_PRODUCTION)
}

func (o Options) ForIntegration(commerceCode, apiKey string) Options {
	return newOptions(commerceCode, apiKey, ENVIRONMENT_INTEGRATION)
}

func (o Options) isProduction() bool {
	return o.Environment == ENVIRONMENT_PRODUCTION
}

func (o Options) GetApiBaseUrl() string {
	if o.isProduction() {
		return BASE_URL_PRODUCTION
	}
	return BASE_URL_INTEGRATION
}
