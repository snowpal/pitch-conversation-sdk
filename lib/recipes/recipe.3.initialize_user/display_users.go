package recipes

import (
	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/endpoints/conversations"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func displayUser(index int, email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info(index+1, ". ", email, " | ", user.JwtToken)
}

func DisplayUsers(userEmails []string) {
	log.Info("## Registered Users")
	for index, userEmail := range userEmails {
		displayUser(index, userEmail)
	}
}

func DisplayMessages(user response.User, conversationId string) {
	log.Info("## Messages for ", user.Email)
	converstion, err := conversations.GetConversation(user.JwtToken, conversationId)
	if err != nil {
		return
	}
	for index, message := range *converstion.Messages {
		log.Info(index+1, ". ", message.MessageText)
	}
}
