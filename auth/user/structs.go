package user

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Ipv4      string `json:"ipv4"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterData struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Ipv4     string `json:"ipv4"`
}

type GoogleData struct {
	AccessToken string `json:"googleToken"`
	Action      string `json:"action"`
}
