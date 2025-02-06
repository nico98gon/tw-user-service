package routers

import (
	"context"
	"errors"
	"testing"
	"user-service/internal/domain"
	"user-service/internal/domain/users"
	"user-service/internal/infrastructure/db"
	"user-service/internal/infrastructure/routers"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister_MissingBody(t *testing.T) {
	ctx := context.Background()
	
	resp := routers.Register(ctx)
	
	assert.Equal(t, 400, resp.Status)
	assert.Contains(t, resp.Message, "No se pudo obtener el body")
}

func TestRegister_InvalidJSON(t *testing.T) {
	ctx := context.WithValue(context.Background(), domain.Key("body"), "{invalid}")
	
	resp := routers.Register(ctx)
	
	assert.Equal(t, 400, resp.Status)
	assert.Contains(t, resp.Message, "Error al parsear el body")
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	body := `{
		"name":"Nicola",
		"last_name":"Tesla",
		"email":"nicolatestla@alterna.com",
		"password":"Alterna>Continua-101"
	}`
	ctx := context.WithValue(context.Background(), domain.Key("body"), body)
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	patches.ApplyFunc(db.UserAlreadyExists, func(string) (users.User, bool, string) {
		return users.User{}, true, ""
	})
	
	resp := routers.Register(ctx)
	
	assert.Equal(t, 400, resp.Status)
	assert.Contains(t, resp.Message, "ya existe")
}

func TestRegister_InsertError(t *testing.T) {
	body := `{
		"name":"Nicola",
		"last_name":"Tesla",
		"email":"nicolatestla@alterna.com",
		"password":"Alterna>Continua-101"
	}`
	ctx := context.WithValue(context.Background(), domain.Key("body"), body)
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	patches.ApplyFunc(db.UserAlreadyExists, func(string) (users.User, bool, string) {
		return users.User{}, false, ""
	})
	patches.ApplyFunc(db.InsertRegister, func(users.User) (string, bool, error) {
		return "", false, errors.New("database error")
	})
	
	resp := routers.Register(ctx)
	
	assert.Equal(t, 400, resp.Status)
	assert.Contains(t, resp.Message, "Error al intentar registrar")
}

func TestRegister_Success(t *testing.T) {
	body := `{
		"name":"Nicola",
		"last_name":"Tesla",
		"email":"nicolatestla@alterna.com",
		"password":"Alterna>Continua-101"
	}`
	ctx := context.WithValue(context.Background(), domain.Key("body"), body)
	
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	patches.ApplyFunc(db.UserAlreadyExists, func(string) (users.User, bool, string) {
		return users.User{}, false, ""
	})
	patches.ApplyFunc(db.InsertRegister, func(users.User) (string, bool, error) {
		return "123", true, nil
	})
	
	resp := routers.Register(ctx)
	
	assert.Equal(t, 200, resp.Status)
	assert.Contains(t, resp.Message, "correctamente")
}