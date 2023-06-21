package conversations

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-conversation-sdk/lib/helpers"
)

type SendMessageReqBody struct {
	MessageText string `json:"messageText"`
}

func SendMessageToAnExistingConversation(jwtToken string, conversationId string, reqBody SendMessageReqBody) (response.Conversation, error) {
	resConversation := response.Conversation{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resConversation, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteConversationsSendMessageToAnExistingConversation, conversationId)
	if err != nil {
		return resConversation, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resConversation, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resConversation, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resConversation, err
	}

	err = json.Unmarshal(body, &resConversation)
	if err != nil {
		return resConversation, err
	}
	return resConversation, nil
}
