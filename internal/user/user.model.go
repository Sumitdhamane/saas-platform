package user

type User struct {
	ID        int64  `json:"id"`
	TenantID  int64  `json:"tenant_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
