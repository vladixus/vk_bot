package main

import (
	"VK_Bot/config"
	"VK_Bot/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-vk-api/vk"
	lp "github.com/go-vk-api/vk/longpoll/user"
)

func main() {
	config.Conf()
	client, err := vk.NewClientWithOptions(
		vk.WithToken(os.Getenv("token")),
		vk.WithHTTPClient(&http.Client{Timeout: time.Second * 30}),
	)
	if err != nil {
		log.Panic(err)
	}

	longpoll, err := lp.NewWithOptions(client, lp.WithMode(lp.ReceiveAttachments))
	if err != nil {
		log.Panic(err)
	}

	stream, err := longpoll.GetUpdatesStream(0)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case update, ok := <-stream.Updates:
			if !ok {
				return
			}

			switch data := update.Data.(type) {
			case *lp.NewMessage:
				switch strings.ToLower(data.Text) {

				case "начать":
					var sentMessageID int64

					keyboard := models.Keyboard{
						OneTime: false,
						Buttons: models.Buttons{
							{
								{
									Action: models.Action{
										Type:    "text",
										Payload: "{\"button\": \"1\"}",
										Label:   "Button 2",
									},
									Color: "positive",
								},
							},
							{
								{
									Action: models.Action{
										Type:    "text",
										Payload: "{\"button\": \"1\"}",
										Label:   "Button 2",
									},
									Color: "positive",
								},
							},
							{
								{
									Action: models.Action{
										Type:    "text",
										Payload: "{\"button\": \"2\"}",
										Label:   "Button 2",
									},
									Color: "positive",
								},
							},
							{
								{
									Action: models.Action{
										Type:    "text",
										Payload: "{\"button\": \"2\"}",
										Label:   "Button 2",
									},
									Color: "positive",
								},
							},
						},
					}

					keyboardJSON, err := json.Marshal(keyboard)
					if err != nil {
						log.Panic(err)
					}

					keyboardString := string(keyboardJSON)

					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "Привет, это чат-бот!",
						"random_id": 0,
						"keyboard":  keyboardString,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				}
			}
		case err, ok := <-stream.Errors:
			if ok {
				log.Panic(err)
			}
		}
	}
}
