package response

type CryptoResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Type  string `json:"type"`
}