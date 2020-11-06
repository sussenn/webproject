package dao

import (
	"webproject/bookstore/model"
	"webproject/bookstore/utils"
)

//查询所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select * from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func AddBook(b *model.Book) error {
	sqlStr := "INSERT INTO `go_stu`.`books` (`title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}
