package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	profiles "github.com/snowpal/pitch-conversation-sdk/lib/endpoints/profile"
)

func RegisterNewUser() (string, error) {
	var err error
	var user response.User

	for i := 1; ; i += 1 {
		email := fmt.Sprintf("apiuser_rec%d_co@yopmail.com", i)
		log.Info(fmt.Sprintf("Register new user with %s", email))
		user, err = recipes.RegisterUser(email)
		if err != nil {
			log.Info(fmt.Sprintf("%s is already registered.", email))
		} else {
			return user.Email, nil
		}
	}
}

func RegisterNewUserAndGetProfile() (response.Profile, error) {
	var profile response.Profile
	newUserEmail, err := RegisterNewUser()
	if err != nil {
		return profile, err
	}
	var user response.User
	user, err = recipes.SignIn(newUserEmail, lib.Password)
	if err != nil {
		return profile, err
	}
	profile, err = profiles.GetUsersProfile(user.JwtToken)
	if err != nil {
		return profile, err
	}
	return profile, nil
}
