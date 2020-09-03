package user

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"oh-my-anime_gf/app/model/user"
)

const (
	USER_SESSION_MARK = "user_info"
)

type SignUpInput struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

func SignUp(data *SignUpInput) error {
	//if e := gvalid.CheckStruct(data, nil); e != nil {
	//	return errors.New(e.FirstString())
	//}
	if data.Nickname == "" {
		data.Nickname = data.Passport
	}
	if CheckPassport(data.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Passport))
	}
	if CheckNickName(data.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data.Nickname))
	}
	var entity *user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}
	entity.CreateTime = gtime.Now()
	if _, err := user.Save(entity); err != nil {
		return err
	}
	return nil
}

//Passport存在则返回true
func CheckPassport(passport string) bool {
	cnt, err := user.FindCount("passport", passport)
	if err != nil {
		return false
	} else {
		return cnt != 0
	}
}

//NickName存在则返回true
func CheckNickName(nickname string) bool {
	cnt, err := user.FindCount("nickname", nickname)
	if err != nil {
		return false
	} else {
		return cnt != 0
	}
}

func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(USER_SESSION_MARK)
}

func SignIn(passport, password string, session *ghttp.Session) error {
	one, err := user.FindOne("passport=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("账号或密码错误")
	}
	return session.Set(USER_SESSION_MARK, one)
}

func SignOut(session *ghttp.Session) error {
	return session.Remove(USER_SESSION_MARK)
}


