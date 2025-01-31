package repository

import (
    "contact-api/internal/model"
    "gorm.io/gorm"
)

type ContactRepository interface {
    Create(contact *model.Contact) error
    FindAll() ([]model.Contact, error)
    FindByID(id uint) (*model.Contact, error)
    Update(contact *model.Contact) error
    Delete(id uint) error
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

func (r *contactRepository) FindAll() ([]model.Contact, error) {
    var contacts []model.Contact
    err := r.db.Find(&contacts).Error
    return contacts, err
}

func (r *contactRepository) FindByID(id uint) (*model.Contact, error) {
    var contact model.Contact
    err := r.db.First(&contact, id).Error
    if err != nil {
        return nil, err
    }
    return &contact, nil
}

func (r *contactRepository) Update(contact *model.Contact) error {
    return r.db.Save(contact).Error
}

func (r *contactRepository) Delete(id uint) error {
    return r.db.Delete(&model.Contact{}, id).Error
}