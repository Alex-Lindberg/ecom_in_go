package customer

type CustomerResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type CustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
