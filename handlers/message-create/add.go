package messageCreate

import (
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sti-discord-bot/databases"
	"sti-discord-bot/databases/models"
	"time"
)

func HandleAdd(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 4 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid arguments for add command.")
		return
	}
	commandObject, eventName, deadline := args[1], args[2], args[3]
	if commandObject != "schedule" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unsupported add command.")
		return
	}

	deadlineTime, err := time.Parse("02/1/2006;15:04", deadline)
	if err != nil {
		log.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	coll := databases.Client.Database("sti_moment").Collection("schedules")
	var userSchedule models.Schedule
	err = coll.FindOne(context.TODO(), bson.D{{"user_id", m.Author.ID}}).Decode(&userSchedule)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result, err := coll.InsertOne(context.TODO(), models.Schedule{
				UserID: m.Author.ID,
				Events: []models.Event{{
					Name:     eventName,
					Deadline: deadlineTime,
				}},
			})
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("New Schedule \"%v\" Inserted\n", result.InsertedID)
			_, _ = s.ChannelMessageSend(m.ChannelID, "New schedule created successfully!")
			return
		}
		log.Println("Error retrieving user schedule:", err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "An error occurred while fetching your schedule.")
		return
	}

	userSchedule.Events = append(userSchedule.Events, models.Event{Name: eventName, Deadline: deadlineTime})
	_, err = coll.UpdateOne(
		context.TODO(),
		bson.D{{"user_id", m.Author.ID}},
		bson.D{{"$set", bson.D{{"events", userSchedule.Events}}}},
	)
	if err != nil {
		log.Println("Error updating schedule:", err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to update your schedule. Please try again.")
		return
	}
	log.Printf("Event \"%v\" added to user %v's schedule.\n", eventName, m.Author.ID)
	_, _ = s.ChannelMessageSend(m.ChannelID, "Event added to your schedule successfully!")
}
