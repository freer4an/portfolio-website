package repository

import (
	"errors"

	"github.com/google/uuid"
)

var (
	Sessions = make(map[string]uuid.UUID)
)

func AddSession(name string, uuid uuid.UUID) error {
	if _, ok := Sessions[name]; ok {
		return errors.New("session already exists")
	}
	Sessions[name] = uuid
	return nil
}

func GetSessionStr(name string) string {
	if session, ok := Sessions[name]; ok {
		return session.String()
	}
	return ""
}

func DeleteSession(name string) {
	if _, ok := Sessions[name]; ok {
		delete(Sessions, name)
	}
}
