package discord

import (
	"github.com/bwmarrin/discordgo"
	"sti-discord-bot/handlers"
)

// InitDiscordClient will initiate the discord client for the Bot application
func InitDiscordClient(authToken string) (*discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + authToken)
	if err != nil {
		return nil, err
	}

	discord.AddHandler(handlers.AddMessage)
	return discord, nil
}
