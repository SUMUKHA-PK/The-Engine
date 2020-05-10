package routing

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/The-Engine/internal/executor"
	"github.com/SUMUKHA-PK/The-Engine/internal/notifications"
	"github.com/SUMUKHA-PK/The-Engine/internal/types"
)

// IncomingQuery decides the best course of action for an incoming request from the client
func IncomingQuery(w http.ResponseWriter, r *http.Request) {

	clientQuery, err := parseHTTPData(r)
	if err != nil {
		fmt.Printf("Cannot parse: %v\n", err)
	}

	title := "Title"
	data := clientQuery // "Data"
	devices := []string{
		"ubuntu",
	}

	key, value, err := parseData(data)
	if err != nil {
		fmt.Printf("Cannot parse incoming data: %v\n", err)
	}
	switch key {
	case "play":
		executor.PlayMedia(value)
	case "test":
		err = notifications.PushNotification(
			notifications.Notification{
				TimeStamp: time.Now().Format("2006-01-02 3:4:5 PM"),
				Title:     title,
				Data:      data,
			},
			devices,
			r.Host,
		)
		if err != nil {
			log.Printf("Notification for %v cannot be pushed\n", data)
		}
	default:
		err = notifications.PushNotification(
			notifications.Notification{
				TimeStamp: time.Now().Format("2006-01-02 3:4:5 PM"),
				Title:     "Unrecognised command",
				Data:      "",
			},
			devices,
			r.Host,
		)
		if err != nil {
			log.Println("Notification for unrecognised command cannot be pushed.")
		}
	}

}

func parseHTTPData(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseData(data string) (key string, value types.Runnable, err error) {
	splitString := strings.Split(data, " ")
	key = splitString[0]
	value.Category = splitString[1]
	for i := 2; i < len(splitString); i++ {
		if value.Category == "TV" {
			if splitString[i] != "S" {
				value.Name += (splitString[i] + " ")
			} else {
				if i+1 < len(splitString) {
					i++
					value.Season = splitString[i]
					if i+2 < len(splitString) {
						i += 2
						value.Episode = splitString[i]
					} else {
						err = errors.New("unexpected end of input")
						return
					}
				} else {
					err = errors.New("unexpected end of input")
					return
				}
			}
		}
	}
	return
}
