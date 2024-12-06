package handlers

import (
	"demo/database"
	"demo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	DB *database.UserDB // promoted field
}

func NewUser(userDB *database.UserDB) *User {
	return &User{DB: userDB}
}

func (u *User) CreateUser(ctx *gin.Context) {

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
	//user.Id = uint(rand.IntN(1000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err = user.Validate()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	user, err = u.DB.Create(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *User) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}

	_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}

	user, err := u.DB.GetByID(uint(_id))

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *User) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}

	_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}

	n, err := u.DB.DeleteByID(uint(_id))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, n)
}

func (u *User) UpdateUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}

	_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
		return
	}
	user := new(models.User)

	err = ctx.Bind(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	user.Id = uint(_id)
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = u.DB.UpdateByID(user)
	//fmt.Println(err.Error())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, user)
}
