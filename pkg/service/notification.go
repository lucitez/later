package service

import (
	"log"

	"github.com/google/uuid"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type Notification struct {
	userService User
}

func NewNotification(userService User) Notification {
	return Notification{userService}
}

type PushMessage struct {
	To    uuid.UUID
	Body  string
	Data  map[string]string
	Title string
}

const (
	// DefaultSound is the default sound for push messages
	DefaultSound = "default"
)

func (s *Notification) SendMessage(message PushMessage) {
	user, err := s.userService.ByID(message.To)

	if err != nil {
		log.Printf("[WARN] error getting user by id %s", err.Error())
		return
	}

	if user == nil {
		log.Println("[WARN] user not found while sending notification")
		return
	}

	if !user.ExpoToken.Valid {
		log.Println("[WARN] user does not have expo token")
		return
	}

	pushToken, err := expo.NewExponentPushToken("ExponentPushToken[" + user.ExpoToken.String + "]")
	if err != nil {
		log.Println("[WARN] malformatted expo token")
		return
	}

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       pushToken,
			Body:     message.Body,
			Data:     message.Data,
			Sound:    DefaultSound,
			Title:    message.Title,
			Priority: expo.DefaultPriority,
		},
	)
	// Check errors
	if err != nil {
		log.Printf("[WARN] Error publishing push notification: %s\n", err.Error())
		return
	}
	// Validate responses
	if validationErr := response.ValidateResponse(); validationErr != nil {
		log.Printf("[WARN] Error validating push notification response: %s\n", validationErr.Error())
	}
}
