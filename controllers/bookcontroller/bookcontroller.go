package bookcontroller

import (
	"encoding/json"
	"errors"
	"golang-crud/config"
	"golang-crud/helper"
	"golang-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Books
	var booksResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&booksResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "List Books", booksResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var books models.Books

	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	//check author
	var author models.Author
	if err := config.DB.First(&author, books.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}
	if err := config.DB.Create(&books).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success create book", nil)

}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Books
	var bookResponse models.BookResponse

	if err := config.DB.Joins("Author").First(&book, id).First(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail book", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Books

	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookPayload models.Books
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author
	if bookPayload.AuthorID != 0 {
		if err := config.DB.First(&author, bookPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author Not Found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&bookPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Book", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Books

	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Book not found", nil)
		return
	}
	helper.Response(w, 200, "Success delete book", nil)
}
