package repositories

import (
	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryDB {
	return UserRepositoryDB{db: db}
}

func (r UserRepositoryDB) GetAll() ([]User, error) {
	var user []User
	result := r.db.Unscoped().Table("users").Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r UserRepositoryDB) Create(user User) (*User, error) {
	result := r.db.Table("users").Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r UserRepositoryDB) FindByCredentials(email, password string) (*User, error) {
	var user *User
	result := r.db.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
