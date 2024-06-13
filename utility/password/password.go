package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func Check(hashPass, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return err == nil
}
