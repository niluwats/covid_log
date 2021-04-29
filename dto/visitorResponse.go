package dto

type NewVisitorResponse struct {
	NIC         string `json:"nic"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AddressCity string `json:"address_city"`
	ContactNo   string `json:"contact_no"`
}
