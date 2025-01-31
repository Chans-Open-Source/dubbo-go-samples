package main

import (
	"context"
	"fmt"
	"strconv"
)

import (
	"github.com/apache/dubbo-go/config"

	"github.com/dubbogo/gost/log"

	perrors "github.com/pkg/errors"
)

func init() {
	config.SetProviderService(new(UserProvider1))
}

type UserProvider1 struct {
}

func (u *UserProvider1) getUser(userID string) (*User, error) {
	if user, ok := userMap[userID]; ok {
		return &user, nil
	}

	return nil, fmt.Errorf("invalid user id:%s", userID)
}

func (u *UserProvider1) GetUser(ctx context.Context, req []interface{}, rsp *User) error {
	var (
		err  error
		user *User
	)

	gxlog.CInfo("req:%#v", req)
	user, err = u.getUser(req[0].(string))
	if err == nil {
		*rsp = *user
		gxlog.CInfo("rsp:%#v", rsp)
	}
	return err
}

func (u *UserProvider1) GetUser0(id string, name string) (User, error) {
	var err error

	gxlog.CInfo("id:%s, name:%s", id, name)
	user, err := u.getUser(id)
	if err != nil {
		return User{}, err
	}
	if user.Name != name {
		return User{}, perrors.New("name is not " + user.Name)
	}
	return *user, err
}

func (u *UserProvider1) GetUser2(ctx context.Context, req []interface{}, rsp *User) error {
	var err error

	gxlog.CInfo("req:%#v", req)
	rsp.ID = strconv.FormatFloat(req[0].(float64), 'f', 0, 64)
	rsp.Sex = Gender(MAN).String()
	return err
}

func (u *UserProvider1) GetUser3() error {
	return nil
}

func (u *UserProvider1) GetUsers(req []interface{}) ([]User, error) {
	return []User{}, nil
}

func (s *UserProvider1) MethodMapper() map[string]string {
	return map[string]string{
		"GetUser2": "getUser",
	}
}

func (u *UserProvider1) Reference() string {
	return "UserProvider1"
}
