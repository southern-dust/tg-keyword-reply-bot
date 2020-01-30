package main

import (
	//"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//"log"
	//"strconv"
	//"strings"
	//"time"
)



// My version for display Fist/Last Name first of a user.
// String displays a simple text version of a user.
//
// It is normally a user's first/last
// name, but falls back to username
// as available.

func NameFirstString(u *tgbotapi.User) string {
	if u.FirstName != "" || u.LastName != ""{
		name := u.FirstName + u.LastName
		return name
	} else if u.UserName != "" {
		return u.UserName
	}
	return ""
}

