package service

import (
	"fmt"
	"time"

	"github.com/eminoz/graceful/model"
)

type UserService struct{}

func (u UserService) UserProcess(user chan model.User) {

	for v := range user {
		fmt.Print(v)
		time.Sleep(time.Second * 2)
	}

}
