package models

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(pe PetError) *errorResponse {
	return &errorResponse{
		Code:    pe.GetCode(),
		Message: pe.GetMessage(),
	}
}
