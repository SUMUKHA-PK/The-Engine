package notifications

import (
	"errors"
	"log"
	"os/exec"
)

// Notification is a basic skeleton for every notification
// Each notification will implement these
type Notification struct {
	TimeStamp string `json:"timestamp"`
	Title     string `json:"title"`
	Data      string `json:"data"`
}

// PushNotification is responsible for pushing notifications across platforms
func PushNotification(notification Notification, onlineDevices []string, requestingDevice string) error {
	for k := range onlineDevices {
		switch onlineDevices[k] {
		case "ubuntu":
			return pushNotificationOnUbuntu(requestingDevice, notification)
		}
	}
	return errors.New("Device not supported")
}

func pushNotificationOnUbuntu(requestingDevice string, notification Notification) error {
	command := "notify-send"
	arg1 := notification.TimeStamp + ": " + notification.Title
	arg2 := notification.Data
	return executeNotifySend(requestingDevice, command, arg1, arg2)
}

func executeNotifySend(requestingDevice string, command string, args ...string) error {
	cmd := exec.Command(command, args[0], args[1])
	log.Printf("%v running %v for %v", requestingDevice, command, args)
	return cmd.Run()
}
