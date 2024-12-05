package handlers

import (
	"demo/models"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	return new(User)
}

func (u *User) Create(ctx *gin.Context) {

	user := new(models.User)

	// bytes, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// }
	// err = json.Unmarshal(bytes, user)
	err := ctx.Bind(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid user data")
		ctx.Abort()
		return
	}
	user.Id = uint(rand.IntN(1000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err = user.Validate()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, user)
}
