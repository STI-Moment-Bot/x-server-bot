package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"sti-discord-bot/databases"
	"sti-discord-bot/discord"
	"syscall"
)

func main() {
	fmt.Println("Tes")
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get the bot token
	botToken := os.Getenv("BOT_TOKEN")
	if len(botToken) == 0 {
		log.Fatal("Bot token not found")
	}

	// Get the database's URI
	dbUri := os.Getenv("MONGODB_URI")
	if dbUri == "" {
		log.Printf("MONGODB_URI not found")
	}

	// Connect to the database
	err = databases.ConnectDB(dbUri)

	if err != nil {
		log.Printf("Error connecting to database: %s", err)
	}

	// init discord client
	session, err := discord.InitDiscordClient(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// open the bot
	err = session.Open()
	if err != nil {
		log.Fatalf("Error when opening connection to Discord: %v", err)
	}
	defer func(session *discordgo.Session) {
		err := session.Close()
		if err != nil {
			log.Fatalf("Error closing session: %v", err)
		}
	}(session)

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	// keep program running until interrupted
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
