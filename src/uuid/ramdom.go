package uuid

import "github.com/google/uuid"

func Uuid() (uuid.UUID, error) {
	return uuid.NewUUID()
}
