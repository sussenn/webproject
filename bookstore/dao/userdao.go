package dao

import (
	"webproject/bookstore/model"
	"webproject/bookstore/utils"
)

//校验登录
func CheckLogin(username string, password string) (*model.User, error) {
	sqlStr := "select * from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		//panic(err.Error())
		//[注意]即便查不到数据,ID也会给一个0 即返回结果user永远不为nil	 &{0 }
		//所以如果查不到,出现err,在这里直接返回,而不是panic. 同时返回nil
		return nil, err
	}
	return u, nil
}

//校验注册名
func CheckRegs(username string) (bool, error) {
	sqlStr := "select * from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		//panic(err.Error())
		return false, err
	}
	return true, nil
}

func SaveUser(username string, password string, email string) error {
	sqlStr := "INSERT INTO `go_stu`.`users` (`username`, `password`, `email`) VALUES (?,?,?)"
	_, sqlErr := utils.Db.Exec(sqlStr, username, password, email)
	if sqlErr != nil {
		//panic(sqlErr.Error())
		return sqlErr
	}
	return nil
}
