package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maksiecluster/booking/pkg/models"
	"github.com/maksiecluster/booking/pkg/utils"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name     string `json:"n"`
	Email    string `json:"em"`
	Password string `json:"p"`
}

func GetTemplate(w http.ResponseWriter, r *http.Request) {

	var url string

	switch r.URL.Path {
	case "/":
		url = "/index.html"
	case "/card":
		url = "/card.html"
	case "/login":
		url = "/login.html"
	case "/signup":
		url = "/registration.html"
	case "/orders":
		url = "/orders.html"
	case "/rent":
		url = "/rent.html"
	default:
		url = "/404.html"
	}

	parsedTemplate, parseErr := template.ParseFiles("../../templates" + url)
	if parseErr != nil {
		fmt.Println("error parsing file", parseErr)
		return
	}
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func Reg(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	user := CreateUser.CreateUser(u.Name, u.Email, u.Password)
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(res)
	if errWrite != nil {
		log.Println("Failed writing HTTP response in CreateUser method")
	}
}

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	newRooms := models.GetAllRooms()
	res, _ := json.Marshal(newRooms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Println("Failed writing HTTP response in GetAllRooms method")
	}
}

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	ID, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	roomDetails, _ := models.GetRoomById(ID)
	res, _ := json.Marshal(roomDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(res)
	if errWrite != nil {
		log.Println("Failed writing HTTP response")
	}
}

func CreateRent(w http.ResponseWriter, r *http.Request) {
	CreateRoom := &models.Room{}
	utils.ParseBody(r, CreateRoom)
	b := CreateRoom.CreateRoom()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Println("Failed writing HTTP response")
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	ID, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	fmt.Println(ID)
	room := models.DeleteRoom(ID)
	res, _ := json.Marshal(room)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, erWrite := w.Write(res)
	if erWrite != nil {
		log.Println("Failed writing HTTP response")
	}
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var updateRoom = &models.Room{}
	utils.ParseBody(r, updateRoom)
	vars := mux.Vars(r)
	roomId := vars["roomId"]
	ID, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	roomDetails, db := models.GetRoomById(ID)
	if updateRoom.Title != "" {
		roomDetails.Title = updateRoom.Title
	}
	if updateRoom.Text != "" {
		roomDetails.Text = updateRoom.Text
	}
	db.Save(&roomDetails)
	res, _ := json.Marshal(roomDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(res)
	if errWrite != nil {
		log.Println("Failed writing HTTP response")
	}
}
