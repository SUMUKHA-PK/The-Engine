package routing

import (
	"log"
	"net/http"
	"time"

	"github.com/SUMUKHA-PK/The-Engine/internal/notifications"
)

// IncomingQuery decides the best course of action for an incoming request from the client
func IncomingQuery(w http.ResponseWriter, r *http.Request) {
	title := "Title"
	data := "Data"
	devices := []string{
		"ubuntu",
	}
	err := notifications.PushNotification(
		notifications.Notification{
			time.Now().Format("2006-01-02 3:4:5 PM"),
			title,
			data,
		},
		devices,
		r.Host,
	)
	// even this must be a push notification
	if err != nil {
		log.Printf("Service could not be completed: %v", err)
	}
}
