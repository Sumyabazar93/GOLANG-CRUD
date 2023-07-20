package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sumyabazar93/go-bookstore/pkg/controllers"
)

func RegisterBookStore(app *fiber.App) {
	app.Use(
		logger.New(),
		recover.New(),
		requestid.New(),
		limiter.New(),
		cors.New(),
		compress.New(),
	)

	app.Post("/book/", controllers.CreateBook)	
	app.Get("/book/", controllers.GetBook)
	app.Get("/book/:bookId", controllers.GetBookById)
	app.Put("/book/:bookId", controllers.UpdateBook)
	app.Delete("/book/:bookId", controllers.DeleteBook)
}
// gorilla/mux 
// var RegisterBookStore = func (router *mux.Router)  {
// 	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
// 	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
// 	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
// 	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
// }
/*
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
*/