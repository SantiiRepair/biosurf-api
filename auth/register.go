package auth

import (
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

type RegisterData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
    var data RegisterData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user := User{
        Email:    data.Email,
        Password: string(passwordHash),
    }
    err = CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}