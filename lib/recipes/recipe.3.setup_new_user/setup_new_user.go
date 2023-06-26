package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-conversation-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func SetupNewUser() {
	log.Info("Objective: Initialize a few user with dynamic email addresses and create group conversation")
	userEmail, err := RegisterNewUser()
	if err != nil {
		return
	}

	log.Info("Register 2 more users to create group conversation")
	var profile2, profile3 response.Profile
	profile2, err = RegisterNewUserAndGetProfile()
	if err != nil {
		return
	}
	profile3, err = RegisterNewUserAndGetProfile()
	if err != nil {
		return
	}

	var conversation response.Conversation
	conversation, err = CreateConversation(userEmail, []string{profile2.Username, profile3.Username})
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("Display messages for the %s user", profile2.Username))
	DisplayMessages(profile2.Email, conversation.ID)

	log.Info(fmt.Sprintf("Send message to group conversation, %s as %s user", conversation.ID, profile2.Username))
	err = SendMessage(profile2.Email, conversation, "hey user1, how are you?")
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("Display messages for the %s user", profile3.Username))
	DisplayMessages(profile3.Email, conversation.ID)

	DisplayUsers([]string{userEmail, profile2.Email, profile3.Email})
}
