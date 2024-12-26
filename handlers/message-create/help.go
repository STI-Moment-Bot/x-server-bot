package messageCreate

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func HandleHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := &discordgo.MessageEmbed{
		Title:       "Help",
		Description: "Welcome To the jungle!\n\nAku Bot STI versi 2.0...",
		Color:       1,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Perintah dipanggil oleh: " + m.Author.Username + " alias " + m.Author.GlobalName,
		},
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		log.Println(err)
	}
}
