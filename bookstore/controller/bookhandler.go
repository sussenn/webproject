package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"webproject/bookstore/dao"
	"webproject/bookstore/model"
)

//查询所有
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

//新增
func AddBooks(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	imgPath := r.PostFormValue("imgPath")
	//string转float64
	fPrice, _ := strconv.ParseFloat(price, 64)
	//10进制的int64
	fSales, _ := strconv.ParseInt(sales, 10, 0)
	fStock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		Title:  title,
		Author: author,
		Price:  fPrice,
		//int64转int
		Sales:   int(fSales),
		Stock:   int(fStock),
		ImgPath: imgPath,
	}
	err := dao.AddBook(book)
	if err != nil {
		fmt.Println(err)
		return
	}
	//新增后,重新查询所有
	GetBooks(w, r)
}
