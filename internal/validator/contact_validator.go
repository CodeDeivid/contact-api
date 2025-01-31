package validator

import (
    "contact-api/internal/model"
    "regexp"
)

const (
    MinNameLength     = 3
    MinPhoneLength    = 8
    EmailRegexPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type ErrorResponse struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

type Validator struct {
    errors []ErrorResponse
}

func NewValidator() *Validator {
    return &Validator{
        errors: make([]ErrorResponse, 0),
    }
}

func (v *Validator) ValidateName(name string) {
    if len(name) < MinNameLength {
        v.errors = append(v.errors, ErrorResponse{
            Field:   "name",
            Message: "O nome deve ter pelo menos 3 caracteres",
        })
    }
}

func (v *Validator) ValidateEmail(email string) {
    emailRegex := regexp.MustCompile(EmailRegexPattern)
    if !emailRegex.MatchString(email) {
        v.errors = append(v.errors, ErrorResponse{
            Field:   "email",
            Message: "Email inválido",
        })
    }
}

func (v *Validator) ValidatePhone(phone string) {
    if len(phone) < MinPhoneLength {
        v.errors = append(v.errors, ErrorResponse{
            Field:   "phone",
            Message: "O telefone deve ter pelo menos 8 caracteres",
        })
    }
}

func (v *Validator) ValidateContact(contact *model.Contact) []ErrorResponse {
    v.errors = make([]ErrorResponse, 0) // Reset errors before validation
    v.ValidateName(contact.Name)
    v.ValidateEmail(contact.Email)
    v.ValidatePhone(contact.Phone)
    return v.errors
}

func (v *Validator) HasErrors() bool {
    return len(v.errors) > 0
}

func (v *Validator) GetErrors() []ErrorResponse {
    return v.errors
}