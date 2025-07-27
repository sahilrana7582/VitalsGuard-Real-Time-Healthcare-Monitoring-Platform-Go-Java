package dto

type NewRoleDTO struct {
	TenantID    string `json:"tenant_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleDTOResponse struct {
	Message string `json:"message"`
}

type AssignRoleResponse struct {
	Message string `json:"message"`
}
