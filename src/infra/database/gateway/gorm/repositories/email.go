package repositories

import (
	"github.com/gomail/src/domain"
	"github.com/gomail/src/infra/database/gateway/gorm/models"
	"gorm.io/gorm"
)

type EmailRepo struct {
	db *gorm.DB
}

func NewEmailRepo(db *gorm.DB) *EmailRepo {
	return &EmailRepo{
		db: db,
	}
}

func (ref *EmailRepo) SaveMail(email domain.Email) error {
	var modelMail models.Email
	modelMail.FromDomain(email)

	return ref.db.Create(&modelMail).Error

}
