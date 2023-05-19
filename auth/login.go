package auth

import (
    "encoding/json"
	bcrypt "golang.org/x/crypto/bcrypt"
    "net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
    var data LoginData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := GetUserByEmail(data.Email)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
    if err != nil {
        http.Error(w, "Incorrect password", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(user)
}