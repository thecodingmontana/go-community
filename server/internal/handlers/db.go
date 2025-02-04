package handlers

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/thecodingmontana/go-community/internal/database/models"
)

type APIConfig struct {
	DB        *models.Queries
	tokenAuth *jwtauth.JWTAuth
}

func NewAPIConfig(queries *models.Queries, tokenAuth *jwtauth.JWTAuth) *APIConfig {
	return &APIConfig{
		DB:        queries,
		tokenAuth: tokenAuth,
	}
}
