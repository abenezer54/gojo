package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	FullName        string     `gorm:"size:100;not null" json:"full_name" binding:"required"`
	Email           string     `gorm:"size:100;unique;not null" json:"email" binding:"required,email"`
	Password        string     `gorm:"not null" json:"-"`
	PhoneNumber     *string    `gorm:"size:15" json:"phone_number,omitempty"`
	ProfileImageURL *string    `gorm:"type:text" json:"profile_image_url,omitempty"`
	Bio             *string    `gorm:"type:text" json:"bio,omitempty"`
	Role            string     `gorm:"size:10;not null" json:"role" binding:"required,oneof=tenant landlord admin"`
	IsVerified      bool       `gorm:"default:false" json:"is_verified"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UserService interface {
	RegisterUser(ctx context.Context, user *SignupRequest) error
}
type UserController interface {
	SignUp(ctx context.Context, user *SignupRequest) error
}

type SignupRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required"`
}
