package profiles

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers"
)

func UpdateUsername(jwtToken string, username string) error {
	route, err := helpers.GetRoute(lib.RouteProfileUpdateUsername, username)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
