package user

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"-"`
}

type LoginData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegisterData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}