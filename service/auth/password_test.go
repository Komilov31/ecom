package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == "password" {
		t.Errorf("Expected to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Errorf("expected passwords to match hash")
	}

	if ComparePasswords(hash, []byte("notPassword")) {
		t.Errorf("expected password to not match hash")
	}
}
