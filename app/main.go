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

	keyboardMain := models.Keyboard{
		OneTime: false,
		Buttons: models.Buttons{
			{
				{
					Action: models.Action{
						Type:    "text",
						Payload: "{\"button\": \"1\"}",
						Label:   "Интситут Мех. и Роб.",
					},
					Color: "secondary",
				},
			},
			{
				{
					Action: models.Action{
						Type:    "text",
						Payload: "{\"button\": \"1\"}",
						Label:   "ИИТиЦТ",
					},
					Color: "secondary",
				},
			},
			{
				{
					Action: models.Action{
						Type:    "text",
						Payload: "{\"button\": \"1\"}",
						Label:   "Институт Соц. Инж.",
					},
					Color: "secondary",
				},
			},
			{
				{
					Action: models.Action{
						Type:    "text",
						Payload: "{\"button\": \"1\"}",
						Label:   "Интситут Эконом. и Менедж.",
					},
					Color: "secondary",
				},
			},
		},
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

					keyboardJSON, err := json.Marshal(keyboardMain)
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
				case "интситут мех. и роб.":
					var sentMessageID int64

					keyboard := models.Keyboard{
						OneTime: false,
						Buttons: models.Buttons{
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "4 курс - мехатрон",
									},
									Color: "primary",
								},
							},
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "Назад",
									},
									Color: "negative",
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
						"message":   "Выбери курс",
						"random_id": 0,
						"keyboard":  keyboardString,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "институт соц. инж.":
					var sentMessageID int64

					keyboard := models.Keyboard{
						OneTime: false,
						Buttons: models.Buttons{
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "4 курс - соц.инж.",
									},
									Color: "primary",
								},
							},
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "Назад",
									},
									Color: "negative",
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
						"message":   "Выбери курс",
						"random_id": 0,
						"keyboard":  keyboardString,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "иитицт":
					var sentMessageID int64

					keyboard := models.Keyboard{
						OneTime: false,
						Buttons: models.Buttons{
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "4 курс - иитицт",
									},
									Color: "primary",
								},
							},
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "Назад",
									},
									Color: "negative",
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
						"message":   "Выбери курс",
						"random_id": 0,
						"keyboard":  keyboardString,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)

				case "интситут эконом. и менедж.":
					var sentMessageID int64

					keyboard := models.Keyboard{
						OneTime: false,
						Buttons: models.Buttons{
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "4 курс - эконом.",
									},
									Color: "primary",
								},
							},
							{
								{
									Action: models.Action{
										Type:  "text",
										Label: "Назад",
									},
									Color: "negative",
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
						"message":   "Выбери курс",
						"random_id": 0,
						"keyboard":  keyboardString,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)

				case "4 курс - мехатрон":
					var sentMessageID int64
					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "https://view.officeapps.live.com/op/view.aspx?src=https://kosygin-rgu.ru/1AppRGU/rguschedule/uploads/ClassSchedules/2023/02/10/202302102143597883.xlsx",
						"random_id": 0,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "4 курс - иитицт":
					var sentMessageID int64
					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "https://view.officeapps.live.com/op/view.aspx?src=https://kosygin-rgu.ru/1AppRGU/rguschedule/uploads/ClassSchedules/2023/04/27/202304271446599461.xlsx",
						"random_id": 0,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "4 курс - соц.инж.":
					var sentMessageID int64
					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "https://view.officeapps.live.com/op/view.aspx?src=https://kosygin-rgu.ru/1AppRGU/rguschedule/uploads/ClassSchedules/2023/03/21/202303211449054581.xlsx",
						"random_id": 0,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "4 курс - эконом.":
					var sentMessageID int64
					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "https://view.officeapps.live.com/op/view.aspx?src=https://kosygin-rgu.ru/1AppRGU/rguschedule/uploads/ClassSchedules/2023/05/03/202305031057183894.xlsx",
						"random_id": 0,
					}, &sentMessageID); err != nil {
						log.Panic(err)
					}

					log.Println(sentMessageID)
				case "назад":
					var sentMessageID int64
					keyboardJSON, err := json.Marshal(keyboardMain)
					if err != nil {
						log.Panic(err)
					}

					keyboardString := string(keyboardJSON)
					if err = client.CallMethod("messages.send", vk.RequestParams{
						"peer_id":   data.PeerID,
						"message":   "Выбери институт",
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
