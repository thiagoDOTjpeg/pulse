package model

import (
	"errors"
	"strings"
	"time"
)

type Url struct {
	Url        string
	CreatedAt time.Time
	UpdatedAt time.Time
	NextVerificationAt time.Time
}

func NewUrl(url string) (Url, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return Url{}, errors.New("Url Inválida")
	}
	nextVerification := time.Now().Add(5 * time.Minute)
	return Url{ Url: url, CreatedAt: time.Now(), UpdatedAt: time.Now(), NextVerificationAt:  nextVerification }, nil
}