package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maksiecluster/booking/pkg/config"
	"github.com/maksiecluster/booking/pkg/utils"
)

var db *gorm.DB

type Room struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}

type User struct {
	gorm.Model
	Name     string `json:"n"`
	Email    string `json:"em"`
	Password string `json:"p"`
	Token    string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Room{})
}

func (u *User) CreateUser(n string, em string, p string) *User {
	token, err := utils.CreateJWT()
	if err != nil {
		panic(err)
	}
	u.Token = token
	u.Name = n
	u.Email = em
	u.Password = p
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func (r *Room) CreateRoom() *Room {
	db.NewRecord(r)
	db.Create(&r)
	return r
}

func GetAllRooms() []Room {
	var Rooms []Room
	db.Find(&Rooms)
	return Rooms
}

func GetRoomById(Id int64) (*Room, *gorm.DB) {
	var getRoom Room
	db := db.Where("ID=?", Id).Find(&getRoom)
	return &getRoom, db
}

func DeleteRoom(ID int64) Room {
	var room Room
	db.Where("ID=?", ID).Delete(room)
	return room
}
