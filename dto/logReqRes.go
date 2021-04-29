package dto

type NewLogRequest struct {
	NIC       string `json:"nic"`
	CompanyId string `json:"company_id"`
}
type NewLogResponse struct {
	NIC     string `json:"nic"`
	LogTime string `json:"log_time"`
	Date    string `json:"date"`
}
type NewLogResponseAll struct {
	Date        string `json:"date" db:"date"`
	LogTime     string `json:"log_time" db:"log_time"`
	NIC         string `json:"nic" db:"nic"`
	FirstName   string `json:"fname" db:"fname"`
	LastName    string `json:"lname" db:"lname"`
	AddressCity string `json:"address_city" db:"address_city"`
	ContactNo   string `json:"contact_no" db:"contact_no"`
}
type NewLogResponseByDate struct {
	Date        string `json:"date" db:"date"`
	NIC         string `json:"nic" db:"nic"`
	LogTime     string `json:"log_time" db:"log_time"`
	FirstName   string `json:"fname" db:"fname"`
	LastName    string `json:"lname" db:"lname"`
	AddressCity string `json:"address_city" db:"address_city"`
	ContactNo   string `json:"contact_no" db:"contact_no"`
}
type NewLogResponseByNic struct {
	Date        string `json:"date" db:"date"`
	LogTime     string `json:"log_time" db:"log_time"`
	CompanyId   string `json:"company_id" db:"company_id"`
	CompanyName string `json:"company_name" db:"name"`
}
