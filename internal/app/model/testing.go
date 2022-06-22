package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "muxammadiyevmansur1996@gmail.com",
		Password: "password",
	}
}
