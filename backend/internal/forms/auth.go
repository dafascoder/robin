package forms

type SignUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type VerifyAccountForm struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
