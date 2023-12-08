package utils

import "github.com/google/uuid"

func InitUUID() uuid.UUID{
	u := uuid.New()
	for i := 0; i < 16; i++ {
		u[i] = 0
	}
	return u
}