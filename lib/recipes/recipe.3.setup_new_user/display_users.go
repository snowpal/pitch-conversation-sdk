package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/endpoints/conversations"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

func displayUser(index int, email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("%d. %s | %s", index+1, email, user.JwtToken))
}

func DisplayUsers(userEmails []string) {
	log.Info("## Registered Users")
	for index, userEmail := range userEmails {
		displayUser(index, userEmail)
	}
}

func DisplayMessages(email string, conversationId string) {
	log.Info(fmt.Sprintf("Sign in as %s", email))
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("## Messages for %s", email))
	converstion, err := conversations.GetConversation(user.JwtToken, conversationId)
	if err != nil {
		return
	}
	for index, message := range *converstion.Messages {
		log.Info(fmt.Sprintf("%d. %s", index+1, message.MessageText))
	}
}
