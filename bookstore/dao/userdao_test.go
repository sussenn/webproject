package dao

import (
	"fmt"
	"testing"
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
	err := SaveUser("marker", "123", "marker@qq.com")
	if err != nil {
		fmt.Println(err)
	}
}
