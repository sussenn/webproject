package dao

import (
	"fmt"
	"testing"
	"webproject/bookstore/model"
)

func TestCheckLogin(t *testing.T) {
	user, err := CheckLogin("marker", "123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func TestCheckRegs(t *testing.T) {
	b, err := CheckRegs("marker")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

func TestSaveUser(t *testing.T) {
	err := SaveUser("admin", "123", "admin@qq.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%d本书:%v\n", k, v)
	}
}

func TestAddBook(t *testing.T) {
	book := &model.Book{
		Title: "笑傲江湖",
	}
	_ = AddBook(book)
}
