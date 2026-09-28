package utils

import (
	"math/rand"
	"strconv"
)

func GenerateResetToken() string {

	num := rand.Intn(900000) + 100000

	return strconv.Itoa(num)
}
