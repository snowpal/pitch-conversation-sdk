package registration

import (
	"net/http"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers"
)

func ActivateUser(userId string) error {
	route, err := helpers.GetRoute(lib.RouteRegistrationActivateUser, userId)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddAppHeaders(req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
