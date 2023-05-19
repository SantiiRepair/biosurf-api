package auth

import (
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

// Define una estructura para los datos del inicio de sesión.
type LoginData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Función para manejar las solicitudes de inicio de sesión.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
    // Parsea los datos del request como un objeto JSON.
    var data LoginData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Busca el usuario en la base de datos.
    user, err := GetUserByEmail(data.Email)
    if err != nil {
        http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
        return
    }

    // Compara la contraseña ingresada con la almacenada en la base de datos.
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
    if err != nil {
        http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
       return
    }

    // Si todo está bien, devuelve el objeto JSON del usuario como respuesta.
    json.NewEncoder(w).Encode(user)
}