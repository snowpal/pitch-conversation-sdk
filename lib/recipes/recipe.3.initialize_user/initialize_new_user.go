package recipes

import (
	"strings"

	"github.com/snowpal/pitch-conversation-sdk/lib"
	"github.com/snowpal/pitch-conversation-sdk/lib/endpoints/conversations"
	"github.com/snowpal/pitch-conversation-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	profiles "github.com/snowpal/pitch-conversation-sdk/lib/endpoints/profile"
)

func InitializeNewUser() {
	log.Info("Objective: Initialize a few user with dynamic email addresses and create group conversation")
	userEmail, err := RegisterNewUser()
	if err != nil {
		return
	}

	log.Info("Register 2 more users to create group conversation")
	var user2Email, user3Email string
	user2Email, err = RegisterNewUser()
	if err != nil {
		return
	}
	user3Email, err = RegisterNewUser()
	if err != nil {
		return
	}

	var user2, user3 response.User
	user2, err = recipes.SignIn(user2Email, lib.Password)
	if err != nil {
		return
	}
	user3, err = recipes.SignIn(user3Email, lib.Password)
	if err != nil {
		return
	}

	var profile2, profile3 response.Profile
	profile2, err = profiles.GetUsersProfile(user2.JwtToken)
	if err != nil {
		return
	}
	profile3, err = profiles.GetUsersProfile(user3.JwtToken)
	if err != nil {
		return
	}

	log.Info("Sign in as ", userEmail)
	var user response.User
	user, err = recipes.SignIn(userEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("Creating a group conversation with ", user2Email, " and ", user3Email)
	var conversation response.Conversation
	conversation, err = conversations.AddPrivateOrGroupConversation(user.JwtToken, conversations.ConversationReqBody{
		MessageText: "hello all",
		Usernames:   strings.Join([]string{profile2.Username, profile3.Username}, ","),
	})
	if err != nil {
		return
	}

	DisplayMessages(user2, conversation.ID)

	log.Info("Sending a message in the grouop conversation as ", user3Email)
	conversation, err = conversations.SendMessageToAnExistingConversation(user2.JwtToken, conversation.ID, conversations.SendMessageReqBody{
		MessageText: "hey user1",
	})
	if err != nil {
		return
	}

	DisplayMessages(user3, conversation.ID)

	DisplayUsers([]string{userEmail, user2Email, user3Email})
}
