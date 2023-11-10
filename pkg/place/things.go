package place

import (
	"github.com/google/uuid"
)

type Thing interface {
	GetID() uuid.UUID
	GetType() ThingType
	GetPlace() *Place
	SetPlace(*Place)
}
