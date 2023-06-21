package user

import (
	"net/http"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers"
)

func DeactivateUserAccount(jwtToken string) error {
	route, err := helpers.GetRoute(lib.RouteUserDeactivateUserAccount)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
