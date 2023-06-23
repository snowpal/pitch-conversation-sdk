package setupnewuser

import (
	"fmt"
	"strings"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/endpoints/conversations"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func CreateConversation(email string, usernames []string) (response.Conversation, error) {
	var conversation response.Conversation

	log.Info(fmt.Sprintf("Sign in as %s", email))
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return conversation, err
	}

	allUsernames := strings.Join(usernames, ",")
	log.Info(fmt.Sprintf("Creating a group conversation with these users: %s", allUsernames))
	conversation, err = conversations.AddPrivateOrGroupConversation(user.JwtToken, conversations.ConversationReqBody{
		MessageText: "hello all",
		Usernames:   allUsernames,
	})
	if err != nil {
		return conversation, err
	}

	return conversation, nil
}

func SendMessage(email string, conversation response.Conversation, message string) error {
	log.Info(fmt.Sprintf("Sign in as %s", email))
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Sending a message in the group conversation as %s", email))
	conversation, err = conversations.SendMessageToAnExistingConversation(
		user.JwtToken,
		conversation.ID,
		conversations.SendMessageReqBody{
			MessageText: message,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
