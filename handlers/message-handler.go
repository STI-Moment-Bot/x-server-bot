package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
	messageCreate "sti-discord-bot/handlers/message-create"
	"strings"
)

func AddMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ensure the message starts with "sti!"
	if !strings.HasPrefix(m.Content, "sti!") {
		return
	}
	args := parseCommandArgumentsPreserveCase(m.Content)
	if len(args) == 0 {
		return
	}

	commands := []CommandHandler{
		{Keyword: getCommandName("ping"), Handler: messageCreate.HandlePing},
		{Keyword: getCommandName("help"), Handler: messageCreate.HandleHelp},
		{Keyword: getCommandName("add"), Handler: messageCreate.HandleAdd},
		{Keyword: getCommandName("find"), Handler: messageCreate.HandleFind},
		{Keyword: getCommandName("media"), Handler: messageCreate.HandleMedia},
	}

	firstWordOfTheMessage := args[0]
	for _, cmd := range commands {
		if firstWordOfTheMessage == cmd.Keyword {
			cmd.Handler(s, m, args)
			return
		}
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "Command not found")
	if err != nil {
		log.Println(err)
	}
}
