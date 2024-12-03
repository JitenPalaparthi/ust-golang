package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"math/rand/v2"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("not implemented other than post"))
		return
	}

	user := new(models.User)

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(bytes, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = user.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user.Id = uint(rand.IntN(1000000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	bytes, err = json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if utils.ChFileWrite != nil {
		utils.ChFileWrite <- []byte(string(bytes) + "\n")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("file writer is not started"))
	}
	// err = utils.WriteToFile("data/users.txt", bytes)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user successfully created"))
}
