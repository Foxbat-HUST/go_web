package handler

import (
	"go_web/api/usecase/user"
	"go_web/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errInvalidPathParam = errors.BadRequestFromStr("invalid path param")

func CreateUser(uc user.CreateUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		input := user.CreateUserInput{}
		if err := ctx.BindJSON(&input); err != nil {
			handleErr(ctx, err)
			return
		}
		output, err := uc.Exec(input)
		if err != nil {
			handleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, output)
	}
}

func UpdateUser(uc user.UpdateUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			handleErr(ctx, errInvalidPathParam)
		}

		input := user.UpdateUserInput{}
		if err := ctx.BindJSON(&input); err != nil {
			handleErr(ctx, err)
			return
		}

		input.ID = uint32(id)
		output, err := uc.Exec(input)
		if err != nil {
			handleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, output)
	}
}

func DeleteUser(uc user.DeleteUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			handleErr(ctx, errInvalidPathParam)
		}

		output, err := uc.Exec(user.DeleteUserInput{
			ID: uint32(id),
		})
		if err != nil {
			handleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, output)
	}
}

func GetUser(uc user.GetUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			handleErr(ctx, errInvalidPathParam)
		}

		output, err := uc.Exec(user.GetUserInput{
			ID: uint32(id),
		})
		if err != nil {
			handleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, output)
	}
}
