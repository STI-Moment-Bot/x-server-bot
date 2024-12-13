package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const prefix string = "sti!"

func getCommandName(command string) string {
	return prefix + command
}

func AddMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// do not respond to himself
	if m.Author.ID == s.State.User.ID {
		return
	}
	firstWordOfTheMessage := strings.Split(m.Content, " ")[0]
	if firstWordOfTheMessage == getCommandName("ping") {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			log.Println(err)
		}
	} else if firstWordOfTheMessage == getCommandName("help") {
		embed := &discordgo.MessageEmbed{
			Title: "Help",
			Description: `Welcome To the jungle!
						
			Aku Bot STI versi 2.0. Ditulis dengan Bahasa Pemrograman Go oleh mas zeev-haydar (https://github.com/zeev-haydar)
							
			Jenis Command yang tersedia untuk saat ini:

			**sti!ping**: Untuk ngetes 
			`,
			Color: 1,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Perintah dipanggil oleh: " + m.Author.Username + " alias " + m.Author.GlobalName,
			},
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			log.Println(err)
		}

	}
}
