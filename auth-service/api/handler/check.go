package handler

import (
	"regexp"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
)

func IsEmailValid(email string) bool {
	rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}

	return true
}

func IsEmailUniq(users *pb.AllUsers, email string) bool {
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Email == email {
			return false
		}
	}

	return true
}

func IsUserNameUniq(users *pb.AllUsers, userName string) bool {
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == userName {
			return false
		}
	}

	return true
}

func IsPhoneValid(phone_number string) bool {
	rxPhone, _ := regexp.MatchString("^\\+998\\d{9}$", phone_number)

	if rxPhone {
		return true
	} else {
		return false
	}
}

func IsNameValid(name string) bool {
	rxName, _ := regexp.MatchString(`^[a-zA-Z]{3,}(?: [a-zA-Z]{3,})?$`, name)

	if rxName {
		return true
	} else {
		return false
	}
}
