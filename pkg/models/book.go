package models

import (
	// "time"

	"github.com/sumyabazar93/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db  *gorm.DB

type Book struct{
	// gorm.model 
	gorm.Model
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}
/*
type Admin struct {
	Firstname         string    `json:"firstname" validate:"required"`
	Lastname          string    `json:"lastname" validate:"required"`
	Phone             string    `json:"phone" validate:"required,len=8"`
	Email             string    `json:"email" validate:"required,email"`
	Hash              string    `json:"-"`
	Salt              string    `json:"-"`
	Gender            string    `json:"gender" validate:"required,oneof=F M"`
	RegistrationNumber string    `json:"registrationNumber"`
	CreatedAtDatetime  time.Time `json:"createdAtDatetime"`
	BirthDate         time.Time `json:"birthDate" validate:"required,lt"`
}
*/

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func(b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}
func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete((book))
	return book
}