package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func main() {
	http.HandleFunc("/api/login", LoginHandler)
	http.HandleFunc("/api/protected", ProtectedHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)

	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	println("체크")
	var user struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// db에서 사용자 정보를 확인 후 넘어가기
	//if user.UserName != "testuser" || user.Password != "testpassword" {
	//	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	//	return
	//}

	// JWT 생성
	// 토큰 인증 만료 시간 24시간으로 설정
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Authorization 헤더에서 JWT 토큰을 추출
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Bearer 토큰 부분을 분리
	tokenString := strings.Split(authHeader, "Bearer ")[1]

	// 토큰을 검증하고 클레임 추출
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil // 서버에서 설정한 비밀 키 사용
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s! You have access to the protected resource.", claims.UserName)
}
