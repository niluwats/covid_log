package dto

type NewCompanyRequest struct {
	CompanyId   string `json:"company_id"`
	CompanyName string `json:"company_name"`
}
type NewCompanyResponse struct {
	CompanyId   string `json:"company_id"`
	CompanyName string `json:"company_name"`
}
