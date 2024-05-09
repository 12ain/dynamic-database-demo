package repository

import (
	"dynamic-database-demo/model"

	"gorm.io/gorm"
	"time"
)

// NewMySQLUserRepository creates a new MySQLUserRepository
func NewMySQLUserRepository(db *gorm.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

// MySQLUserRepository implements UserRepository interface for MySQL
type MySQLUserRepository struct {
	db *gorm.DB
}

func (m *MySQLUserRepository) CreateUser(user model.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.ID = uint(now.UnixNano())
	result := m.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MySQLUserRepository) GetUser(userID string) (model.User, error) {
	var user model.User
	result := m.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (m *MySQLUserRepository) UpdateUser(user model.User) error {
	result := m.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *MySQLUserRepository) DeleteUser(userID string) error {
	result := m.db.Where("id = ?", userID).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
