package service

import (
	"errors"
	"fmt"
	"gjob-admin/pkg/utils"
)

var Login login

type login struct{}

// 验证账号密码
func (l *login) Auth(username, password string) (err error) {
	if username == utils.Config.GetString("User.adminuser") && password == utils.Config.GetString("User.adminpwd") {
		return nil
	}
	fmt.Println("登录失败, 用户名或密码错误")
	return errors.New("登录失败, 用户名或密码错误")
}
