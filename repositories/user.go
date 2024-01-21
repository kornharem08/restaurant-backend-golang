package repositories

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	Create(User) (*User, error)
	FindByCredentials(email string, password string) (*User, error)
}
