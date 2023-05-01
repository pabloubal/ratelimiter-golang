package dto

type CreateEndpointDto struct {
	Path   string `json:"path"`
	Url    string `json:"url"`
	Limit  int    `json:"limit"`
	Method string `json:"method"`
}
