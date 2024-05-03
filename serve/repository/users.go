package repository

import (
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
	"gorm.io/gorm"
)

type UsersRepository interface {
	SaveOrUpdate(user *models.User) error
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id string) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{db: conn.DB()}
}

func (r *usersRepository) SaveOrUpdate(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *usersRepository) GetById(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *usersRepository) Delete(id string) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
