package entity

type UserRepository struct {
	UserID    string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
