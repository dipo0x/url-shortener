package types

type IURL struct {
    URL       string `json:"url"`
    ExpiresAt string `json:"expiresAt"`
}
type IDeletePayload struct {
	DocumentID string `json:"document_id"`
}
