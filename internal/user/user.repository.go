package user

import (
	"github.com/sumitdhamane/saas-platform/internal/database"
)

func CreateUser(
	tenantID int64,
	firstName string,
	lastName string,
	email string,
	password string,
) (int64, error) {

	query := `
	INSERT INTO users(
		tenant_id,
		first_name,
		last_name,
		email,
		password
	)
	VALUES(?,?,?,?,?)
	`

	result, err := database.DB.Exec(
		query,
		tenantID,
		firstName,
		lastName,
		email,
		password,
	)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
