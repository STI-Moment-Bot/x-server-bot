package models

import "time"

type Schedule struct {
	UserID string  `bson:"user_id,omitempty"`
	Events []Event `bson:"events,omitempty"`
}

type Event struct {
	Name     string    `bson:"name"`
	Deadline time.Time `bson:"deadline,omitempty"`
}

func FetchSchedule(userID string) (*Schedule, error) {
	return &Schedule{}, nil
}

func AddEventOnSchedule(event Event) error {
	return nil
}
