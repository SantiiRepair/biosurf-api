package user

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	GoogleAccount   string `json:"google_account"`
	FacebookAccount string `json:"facebook_account"`
	GoogleID        string `json:"google_id"`
	FacebookID      string `json:"facebook_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Ipv4            string `json:"ipv4"`
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
	JWTDataUser string `json:"googleToken"`
	Action      string `json:"action"`
}
