package domain

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doel"`
	UserName  string `json:"username" example:"jhon_doel"`
	Email     string `json:"email" example:"user@email.com"`
	Enabled   bool   `json:"enabled"`
}

type UserLogin struct {
	Username string `json:"username" example:"jhondoel"`
	Password string `json:"password" example:"s3cr3t3"`
}

type Token struct {
	Token     string `json:"token" example:"asdfas-asdfasd-asdf-asdf-asdf"`
	ExpiredIn int    `json:"expires_in"`
}

type CreateUser struct {
	FirstName string `json:"first_name" example:"Jhon"`
	LastName  string `json:"last_name" example:"Doel"`
	Email     string `json:"email" example:"jhon.doel@example.com"`
	UserName  string `json:"username" example:"jhondoel1995"`
	Password  string `json:"password" example:"s3cr3t3"`
}
