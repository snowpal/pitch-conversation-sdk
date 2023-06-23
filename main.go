package main

import (
	"os"
	"strconv"

	"github.com/snowpal/pitch-conversation-sdk/lib/config"
	"github.com/snowpal/pitch-conversation-sdk/lib/recipes"

	log "github.com/sirupsen/logrus"
	setupnewuser "github.com/snowpal/pitch-conversation-sdk/lib/recipes/recipe.3.setup_new_user"
)

func main() {
	var err error
	if config.Files, err = config.InitConfigFiles(); err != nil {
		log.Fatal(err.Error())
		return
	}

	var recipeID int
	recipeIDInEnv := os.Getenv("RECIPE_ID")
	if len(recipeIDInEnv) == 0 {
		recipeID = 1
	} else {
		recipeID, err = strconv.Atoi(recipeIDInEnv)
		if err != nil {
			recipeID = 1
		}
	}

	switch recipeID {
	case 1:
		log.Info("Run Recipe1")
		recipes.RegisterFewUsers()
		break
	case 2:
		log.Info("Run Recipe2")
		recipes.CreatePrivateConversation()
		break
	case 3:
		log.Info("Run Recipe3")
		setupnewuser.SetupNewUser()
	default:
		log.Info("pick a specific recipe to run")
	}
}
