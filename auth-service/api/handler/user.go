package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	token "github.com/dilshodforever/nasiya-savdo/api/token"
	"github.com/dilshodforever/nasiya-savdo/config"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

type changePass struct {
	CurrentPassword string
	NewPassword     string
}

type resetPass struct {
	ResetToken  string
	NewPassword string
}

// RegisterUser handles the creation of a new user
// @Summary Register User
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body pb.UserReq true "Create"
// @Success 201 {string} string "Create Successful"
// @Failure 400 {string} string "Error while creating user"
// @Router /user/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	user := pb.UserReq{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	req, err := json.Marshal(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	h.producer.ProduceMessages("user-create", req)

	ctx.JSON(http.StatusCreated, "Success!!!")
}

// UpdateUser handles updating an existing user
// @Summary Update User
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body pb.User true "Update"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	user := pb.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.User.Update(ctx, &user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// DeleteUser handles the deletion of a user
// @Summary Delete User
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	_, err := h.User.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// GetAllUser handles retrieving all users
// @Summary Get All Users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query pb.UserFilter true "Query parameter"
// @Success 200 {object} pb.AllUsers "Get All Successful"
// @Failure 400 {string} string "Error while retrieving users"
// @Router /user/getall [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	filter := pb.UserFilter{}
	err := ctx.BindQuery(&filter)
	if err != nil {
		ctx.JSON(400, err.Error())
	}

	res, err := h.User.GetAll(ctx, &filter)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// GetByIdUser handles retrieving a user by ID
// @Summary Get User By ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} pb.User "Get By ID Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Router /user/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	res, err := h.User.GetById(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// LoginUser handles user login
// @Summary Login User
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Create body pb.UserLogin true "Create"
// @Success 200 {object} string "Login Successful"
// @Failure 400 {string} string "Error while logging in"
// @Router /user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := pb.UserLogin{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	res, err := h.User.Login(ctx, &user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	t := token.GenereteJWTToken(res)

	h.redis.Set(res.Id, t.RefreshToken, 30*24*time.Hour)

	ctx.JSON(200, t)
}

// ChangePassword handles changing user password
// @Summary Change Password
// @Description Change user password
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ChangePass body changePass true "Change Password"
// @Success 200 {string} string "Password Changed Successfully"
// @Failure 400 {string} string "Error while changing password"
// @Router /user/change-password [post]
func (h *Handler) ChangePassword(ctx *gin.Context) {
	changePas := changePass{}
	err := ctx.BindJSON(&changePas)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	changePass := pb.ChangePass{CurrentPassword: changePas.CurrentPassword, NewPassword: changePas.NewPassword}
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	changePass.Id = id

	_, err = h.User.ChangePassword(ctx, &changePass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Password Changed Successfully")
}

// ForgotPassword handles initiating the forgot password process
// @Summary Forgot Password
// @Description Initiate forgot password process
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ForgotPass body pb.ForgotPass true "Forgot Password"
// @Success 200 {string} string "Forgot Password Initiated Successfully"
// @Failure 400 {string} string "Error while initiating forgot password"
// @Router /user/forgot-password [post]
func (h *Handler) ForgotPassword(ctx *gin.Context) {
	forgotPass := pb.ForgotPass{}
	err := ctx.BindJSON(&forgotPass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	f := rand.Intn(899999) + 100000
	err = h.redis.SaveToken(forgotPass.Email, fmt.Sprintf("%d", f), time.Minute*2)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.User.ForgotPassword(ctx, &forgotPass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Email message has been sent")
}

// ResetPassword handles resetting the user password
// @Summary Reset Password
// @Description Reset user password
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param ResetPass body resetPass true "Reset Password"
// @Success 200 {string} string "Password Reset Successfully"
// @Failure 400 {string} string "Error while resetting password"
// @Router /user/reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	resetPas := resetPass{}
	err := ctx.BindJSON(&resetPas)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resetPass := pb.ResetPass{ResetToken: resetPas.ResetToken, NewPassword: resetPas.NewPassword}
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	resetPass.Id = id

	email, _ := token.GetEmailFromToken(ctx.Request, &cnf)
	e, err := h.redis.Get(email)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	if e != resetPass.ResetToken {
		ctx.JSON(400, "Invalid reset-password")
		return
	}

	_, err = h.User.ResetPassword(ctx, &resetPass)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Password Reset Successfully")
}

// GetProfil handles retrieving a user Profil
// @Summary Get User Profil
// @Description Get a user Profil
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.User "Get Profil Successful"
// @Failure 400 {string} string "Error while retrieving user"
// @Router /user/get_profil [get]
func (h *Handler) GetProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	res, err := h.User.GetById(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// UpdateUser handles updating an existing user
// @Summary Update Profil
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Update body pb.User true "Update"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Error while updating user"
// @Router /user/update_profil [put]
func (h *Handler) UpdateProfil(ctx *gin.Context) {
	user := pb.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	cnf := config.Load()
	user.Id, _ = token.GetIdFromToken(ctx.Request, &cnf)
	_, err = h.User.Update(ctx, &user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// DeleteProfil handles the deletion of a Profil
// @Summary Delete Profil
// @Description Delete an existing Profil
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while deleting user"
// @Router /user/delete_profil [delete]
func (h *Handler) DeleteProfil(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)

	_, err := h.User.Delete(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// RefreshToekn handles the deletion of a Token
// @Summary refresh Toekn
// @Description refresh an existing Token
// @Tags Auth
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "refresh Successful"
// @Failure 400 {string} string "Error while refreshed token"
// @Router /user/refresh-token [get]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	cnf := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &cnf)
	tok, err := h.redis.Get(id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = token.ExtractClaim(&cnf, tok)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	user, err := h.User.GetById(ctx, &pb.ById{Id: id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	access := token.GenereteAccsessJWTToken(user).AccessToken

	ctx.Request.Header.Set("Authorization", access)

	ctx.JSON(200, gin.H{"new access token: ": access})
}
