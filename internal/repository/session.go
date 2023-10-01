package repository

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	Sessions = make(map[string]uuid.UUID)
)

func AddSession(name string, uuid uuid.UUID) {
	Sessions[name] = uuid
}

func GetSessionStr(name string) (string, error) {
	if session, ok := Sessions[name]; ok {
		return session.String(), nil
	}
	return "", fmt.Errorf("Session not found")
}

func DeleteSession(name string) {
	if _, ok := Sessions[name]; ok {
		delete(Sessions, name)
	}
}
