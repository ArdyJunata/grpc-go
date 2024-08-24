package auth

import (
	"time"

	"github.com/ArdyJunata/grpc-go/utility"
)

type Auth struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *Auth) encryptPassword(salt int) (err error) {
	pass, err := utility.Encrypt(a.Password, salt)
	if err != nil {
		return err
	}

	a.Password = pass

	return
}

func (a Auth) VerifyPlainPassword(plain string) error {
	err := utility.Verify(plain, a.Password)
	return err
}

func (a Auth) GenerateJWT(secret string, duration int16) (token string, err error) {
	return utility.GenerateJWT(utility.JWTPayload{
		Id: a.Id,
	}, secret, duration)
}
