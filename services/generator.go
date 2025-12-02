package services

import (
	"fmt"
	"math/rand"
	"time"
)

const aleph = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+?"

// generates a password of a particular length based 
// on the input parameter
func GeneratePassword(n int) string {
	
	rand.Seed(time.Now().UnixNano())
	pass := make([]byte, n)
	
	for i := range pass {
		pass[i] = alep[rand.Intn(len(aleph))]
	}
	
	return string(pass)
}

