package model

type Hash struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}

func (h Hash) ToResponse() HashResponse {
	return HashResponse{
		ID:    h.ID,
		Value: h.Value,
	}
}

type HashResponse struct {
	ID    uint   `json:"id"`
	Value string `json:"value"`
}
