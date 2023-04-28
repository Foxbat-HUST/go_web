package handler

import (
	"go_web/api/usecase/user"
	"go_web/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errInvalidPathParam = errors.BadRequestFromStr("invalid path param")

func createUser(uc user.CreateUser) func(ctx *gin.Context) {
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

func updateUser(uc user.UpdateUser) func(ctx *gin.Context) {
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

func deleteUser(uc user.DeleteUser) func(ctx *gin.Context) {
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

func getUser(uc user.GetUser) func(ctx *gin.Context) {
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

func listUser(uc user.ListUser) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var queryParam struct {
			PageIndex   int `form:"p"`
			ItemPerPage int `form:"l"`
		}

		err := ctx.BindQuery(&queryParam)
		if err != nil {
			handleErr(ctx, errors.BadRequest(err))
			return
		}

		output, err := uc.Exec(user.ListUserInput{
			PageIndex:   queryParam.PageIndex,
			ItemPerPage: queryParam.ItemPerPage,
		})

		if err != nil {
			handleErr(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, output)
	}
}
