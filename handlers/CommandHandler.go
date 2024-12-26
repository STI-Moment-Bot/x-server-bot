package handlers

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const prefix string = "sti!"

type CommandHandler struct {
	Keyword string
	Handler func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

func getCommandName(command string) string {
	return prefix + command
}

func parseCommandArgumentsPreserveCase(content string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false

	for _, char := range content {
		switch {
		case char == '"' && inQuotes:
			inQuotes = false
			args = append(args, currentArg.String())
			currentArg.Reset()
		case char == '"' && !inQuotes:
			inQuotes = true
		case char == ' ' && !inQuotes:
			if currentArg.Len() > 0 {
				args = append(args, strings.ToLower(currentArg.String()))
				currentArg.Reset()
			}
		default:
			currentArg.WriteRune(char)
		}
	}

	if currentArg.Len() > 0 {
		if inQuotes {
			args = append(args, currentArg.String())
		} else {
			args = append(args, strings.ToLower(currentArg.String()))
		}
	}

	return args
}
