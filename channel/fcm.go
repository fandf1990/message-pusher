package channel

import (
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"message-pusher/model"
)

func SendFirebaseCloudMessage(message *model.Message, user *model.User, channel_ *model.Channel) error {
	//https://studygolang.com/articles/19998

	//opt := option.WithCredentialsFile("C:\\firebase-adminsdk.json")

	if channel_.Secret == "" || channel_.AccountId == "" {
		return errors.New("未配置 Google FCM 消息推送方式")
	}
	var serviceAccountKey = []byte(channel_.Secret)
	opt := option.WithCredentialsJSON(serviceAccountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	notification := &messaging.Notification{
		Title: message.Title,
		Body:  message.Content,
	}
	pushMsg := &messaging.Message{
		Notification: notification,
		Token:        channel_.AccountId,
	}

	response, err := client.Send(ctx, pushMsg)
	if err != nil {
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return nil
}
