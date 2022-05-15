package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	UserId string
	Name   string
}

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代
func GetUserById(id string) (User, error) {
	var user User

	err := sql.ErrNoRows

	if err != nil {
		// 应该包装一下底层错误，解耦
		return errors.Wrap(err, "user not found")
	}

	return user, nil
}
