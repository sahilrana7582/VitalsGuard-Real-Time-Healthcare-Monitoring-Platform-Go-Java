package dto

type SignUpRequest struct {
	TenantID string `json:"tenant_id" validate:"required,uuid4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}
