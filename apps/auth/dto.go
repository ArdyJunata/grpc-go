package auth

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d registerRequest) ParseToModel() Auth {
	return Auth{
		Username: d.Username,
		Password: d.Password,
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d loginRequest) ParseToModel() Auth {
	return Auth{
		Username: d.Username,
		Password: d.Password,
	}
}
