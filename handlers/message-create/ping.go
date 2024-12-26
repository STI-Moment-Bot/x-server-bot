package messageCreate

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func HandlePing(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Println(err)
	}
}
