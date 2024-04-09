package entities

type LoginDetails struct {
	AuthToken string
	Username  string
	Mail      string
}

// DTO (Secure version of LoginDetails)
type UserDetails struct {
	Username string
	Mail     string
}
