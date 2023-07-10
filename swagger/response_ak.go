package swagger

type GetAkResponse struct {
	RequestId string        `json:"request_id,omitempty"`
	Code      string        `json:"code,omitempty"`
	Message   string        `json:"message,omitempty"`
	Data      map[string]Ak `json:"data,omitempty"`
}