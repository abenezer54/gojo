package controller

import (
	"net/http"

	"github.com/abenezer54/gojo/backend/user-service/internal/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService model.UserService
}

func NewUserController(userService model.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

type SignUpRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=tenant landlord admin"`
}

func (c *UserController) SignUp(ctx *gin.Context) {
	var req SignUpRequest

	// Bind and validate JSON request.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map to the service layer's model.
	signupModel := &model.SignupRequest{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	// Call the user service to perform registration.
	err := c.userService.RegisterUser(ctx.Request.Context(), signupModel)
	if err != nil {
		// Handle known "email already in use" error.
		if err.Error() == "email already in use" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// Handle all other internal errors.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Return success response.
	ctx.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}
