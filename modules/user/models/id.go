package models

import (
	"math/rand/v2"
)

type ID int64

func (id ID) generateRandomId() int64 {
	randomInt := rand.Int64N(1000000)
	return randomInt
}
