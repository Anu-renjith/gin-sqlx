package entity

type User struct {
	ID    uint   `db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
