package messageCreate

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sti-discord-bot/databases"
	"sti-discord-bot/databases/models"
)

func HandleFind(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid arguments for find command.")
		return
	}
	commandObject := args[1]
	if commandObject != "schedule" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unsupported find command.")
		return
	}

	coll := databases.Client.Database("sti_moment").Collection("schedules")
	var userSchedule models.Schedule
	err := coll.FindOne(context.TODO(), bson.D{{"user_id", m.Author.ID}}).Decode(&userSchedule)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, _ = s.ChannelMessageSend(m.ChannelID, "You have no schedules")
			return
		}
		log.Printf("ERROR: %v", err.Error())
		_, _ = s.ChannelMessageSend(m.ChannelID, "An error occurred while fetching your schedule.")
		return
	}

	var willSendMessage = fmt.Sprintf("## %v's Schedule\n\n", m.Author.Username)
	for index, event := range userSchedule.Events {
		willSendMessage += fmt.Sprintf("**%v. %v**Deadline: %v %v %v %v.%v\n",
			index+1, event.Name, event.Deadline.Day(), event.Deadline.Month().String(), event.Deadline.Year(),
			fmt.Sprintf("%02d", event.Deadline.Hour()), fmt.Sprintf("%02d", event.Deadline.Minute()))
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, willSendMessage)
}
