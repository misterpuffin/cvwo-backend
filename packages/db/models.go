package db

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"server/packages/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint `json:"id,omitempty" gorm:"primaryKey"`
	Password  string `json:"-"`
	Email     string `json:"email,omitempty" gorm:"unique,not null;default:null"`
	Name      string `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"-"`
}

type JWTToken struct {
	Token string `json:"token"`
	Name string `json:"name"`
}

func (u User) HashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

func (u User) CheckPasswordHash(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}


func (u User) GenerateJWT() (JWTToken, error) {
	signingKey := []byte(config.Config[config.JWT_KEY])
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1 * 1).Unix(),
		"user_id": int(u.ID),
		"name": u.Name,
		"email": u.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString, u.Name}, err
}

type Task struct {
	ID        uint `json:"id,omitempty" gorm:"primaryKey"`
	UserID	  string `json:"user_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Tags	  []string `json:"tags,omitempty" gorm:"-"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"-"`
}

type Tag struct {
	Name	string `json:"name,omitempty"`
	UserID	string `json:"user_id,omitempty"`
	TaskID	string `json:"task_id,omitempty"`
}