package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/mail"
	"time"

	"github.com/crit/critical-go/storage"
)

var cache = storage.Local()

// Person represents a registrant.
type Person struct {
	Id    int       `json:"-"`
	Name  string    `json:"name" sql:"not null;unique"`
	Email string    `json:"-"  sql:"not null"`
	Date  time.Time `json:"date"  sql:"not null;DEFAULT:current_timestamp"`
}

func (Person) TableName() string {
	return "people"
}

func MigratePerson() {
	if db.HasTable("people") {
		return
	}

	err := db.AutoMigrate(&Person{}).Error

	if err != nil {
		log.Println(err.Error())
	}
}

func PersonList() []Person {
	list := make([]Person, 0)

	data := cache.Get("people")

	if err := json.Unmarshal(data, &list); len(data) > 0 && err == nil {
		return list
	}

	db.Order("date desc").Find(&list)

	if len(list) > 0 {
		data, _ := json.Marshal(list)
		cache.Put("people", data)
	}

	return list
}

func PersonCreate(name, email string) error {
	if name == "" {
		return errors.New("Missing name from post.")
	}

	if email != "" && notEmail(email) {
		return errors.New("Invalid email format.")
	}

	person := Person{Name: name, Email: email, Date: time.Now()}
	err := db.Create(&person).Error

	if err == nil {
		cache.Delete("people")
	}

	return err
}

func notEmail(in string) bool {
	_, err := mail.ParseAddress(in)

	return err != nil
}
