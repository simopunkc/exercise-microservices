package util

import (
	"golang.org/x/crypto/bcrypt"
)

type UtilBcrypt struct {
}

func NewUtilBcrypt() *UtilBcrypt {
	return &UtilBcrypt{}
}

func (ub UtilBcrypt) CheckIfPasswordHashIsEqual(repoPassword []byte, paramPassword []byte) error {
	return bcrypt.CompareHashAndPassword(repoPassword, paramPassword)
}

func (ub UtilBcrypt) GenerateHashFromPlainPassword(paramPassword []byte) (string, error) {
	result, err := bcrypt.GenerateFromPassword(paramPassword, bcrypt.DefaultCost)
	return string(result), err
}
