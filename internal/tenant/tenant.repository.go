package tenant

import (
	"github.com/sumitdhamane/saas-platform/internal/database"
)

func CreateTenant(name string) (int64, error) {
	query := `
	INSERT INTO tenants(name)
	VALUES(?)
	`

	result, err := database.DB.Exec(query, name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
