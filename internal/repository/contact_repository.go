package repository

import (
	"contact-api/internal/model"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *model.Contact) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

func (r *contactRepository) Create(contact *model.Contact) error {
	return r.db.Create(contact).Error
}