package dao

import (
	"fmt"
	"github.com/fu-hui/am/src/model"
)

// create user
func CreateUser(user model.User) error {
	createUserSql := "insert into user(username,password) values(?,?)"
	_, err := AmDb.Exec(createUserSql, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("insert user failed, err:%v", err)
	}
	return nil
}

// select user by user name
func QueryUserByUsername(username string) (*model.User, error) {
	createUserSql := "select username, password from user where username = ?"
	rows, err := AmDb.Query(createUserSql, username)
	if err != nil {
		return nil, fmt.Errorf("select user failed, err:%v", err)
	}
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if !rows.Next() {
		return nil, nil
	}

	var user model.User
	err = rows.Scan(&user.UserName, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("scan user rows failed, err:%v", err)
	}

	return &user, nil
}
