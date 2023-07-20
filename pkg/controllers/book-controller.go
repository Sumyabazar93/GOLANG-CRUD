package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sumyabazar93/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(c *fiber.Ctx) error{
	newBooks:= models.GetAllBooks()
	res, err:= json.Marshal(newBooks)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}
	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).Send(res)
}

func GetBookById(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}
	bookDetails, _ := models.GetBookById(ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Book nor found")
	}

	res, err := json.Marshal(bookDetails)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}
	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).Send(res)
}

func CreateBook(c *fiber.Ctx) error{
	CreateBook := &models.Book{}

	if err := c.BodyParser(CreateBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	b:=CreateBook.CreateBook()
	res, err:= json.Marshal(b)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}
	return c.Status(fiber.StatusOK).Send(res)
}

func DeleteBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}
	book := models.DeleteBook(ID)

	res, err := json.Marshal(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}
	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).Send(res)
}
func UpdateBook(c *fiber.Ctx) error {
	var updateBook = &models.Book{}

	if err := c.BodyParser(updateBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	bookId := c.Params("bookId")
	
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	
	db.Save(&bookDetails)
	res, err:= json.Marshal(bookDetails)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}

	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).Send(res)
}
/*
 Gorilla mux 

 import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
 )
func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= models.GetAllBooks()
	res, _:= json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r, CreateBook)
	b:=CreateBook.CreateBook()
	res, _:= json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsin")
	}
	bookDetails, db:= models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	
	db.Save(&bookDetails)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
*/