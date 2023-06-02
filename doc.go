/*
Check out the recipes to get a sense of what the underlying API has to offer. Here's a simple example:

	package main

	import (
		log "github.com/sirupsen/logrus"
		"github.com/snowpal/pitch-building-blocks-sdk/lib/config"
		"github.com/snowpal/pitch-building-blocks-sdk/lib/recipes"
	)

	func main() {
		var err error
		if config.Files, err = config.InitConfigFiles(); err != nil {
			log.Fatal(err.Error())
			return
		}

		recipeID := 1
		switch recipeID {
		case 1:
			log.Info("Run Recipe1")
			recipes.RegisterFewUsers()
			break
		case 2:
			log.Info("Run Recipe2")
			recipes.CreatePrivateConversation()
			break
		default:
			log.Info("pick a specific recipe to run")
		}
	}

For a full guide, visit https://developers.snowpal.com
*/
package main
