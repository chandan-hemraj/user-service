package client

type ClientInterface interface {
	MakeRequest(url, method string, payload []byte, headers map[string]string) (map[string]interface{}, string, error)
}
