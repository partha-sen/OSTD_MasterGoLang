package service

import (
	"github.com/partha-sen/ostd/authServer/dao"
	"github.com/partha-sen/ostd/authServer/model"
	"github.com/partha-sen/ostd/authServer/token"
	"github.com/pkg/errors"
)

func IssueToken(user model.User) (string, error) {
	if usr, err := dao.ValidateAndGetUser(user); err != nil {
		return "", errors.Wrap(err, "Can't issue token")
	} else {
		token, err := token.CreateToken(usr.Username, usr.Role)
		return token, err
	}
}
