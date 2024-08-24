package utility

import "golang.org/x/crypto/bcrypt"

func Encrypt(plain string, salt int) (encrypted string, err error) {
	encryptedByte, err := bcrypt.GenerateFromPassword([]byte(plain), salt)
	if err != nil {
		return
	}

	encrypted = string(encryptedByte)
	return
}

func Verify(plain string, encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain))
}
