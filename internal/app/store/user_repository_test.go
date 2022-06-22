package store_test

import (
	"testing"

	"github.com/Muhammadiyev/rest-api/internal/app/model"
	"github.com/Muhammadiyev/rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, testdown := store.TestStore(t, databaseURL)
	defer testdown("users")
	u, err := s.User().Create(&model.User{
		Email: "muxammadiyevmansur1996@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, testdown := store.TestStore(t, databaseURL)
	defer testdown("users")

	email := "muxammadiyevmansur1996@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "muxammadiyevmansur1996@gmail.com",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
