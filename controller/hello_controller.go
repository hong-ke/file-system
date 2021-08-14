package controller

import (
	"filesystem/dig"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get-hello
// @Tags hello
// @Description
// @Accept json
// @Param Authorization header string true "Header Auth Access Token of User"
// @Param id path string true "群组id"
// @Success  204
// @Router /v1/{hello}  [get]
func HelloController(ctx *gin.Context) {
	lock, err := dig.NewRedisLock()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	tryLock, err := lock.TryLock(lock.Prefix + "lock")
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	if tryLock == nil {
		ctx.JSON(http.StatusOK, "nolock")
		return
	}
	ctx.JSON(http.StatusOK, "get")
}

func Hello2Controller(ctx *gin.Context) {
	service, err := dig.NewHelloService()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	//id := uuid.NewV4()
	//u := &entity.User{
	//	ID:       id.String(),
	//	Name:     "hk",
	//	Password: "123456",
	//}
	get, err := service.Get()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	lock, err := dig.NewRedisLock()
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	lock.TryLock(lock.Prefix + "lock")
	ctx.JSON(http.StatusOK, get)
}
