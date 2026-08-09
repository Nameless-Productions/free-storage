package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("No inputs provided")
	}


	if os.Getenv("TOKEN") == "" || os.Getenv("CHANNEL_ID") == "" {
		log.Fatal("No TOKEN in env")
	}

	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	user, err := session.User("@me");
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bot: %s#%s \n", user.Username, user.Discriminator)

	if existsInDB(args[0]) {
		msgIDs := getMessageIDs(args[0])

		fullFile := []byte{}

		for i, msgID := range msgIDs {
			msg, err := session.ChannelMessage(os.Getenv("CHANNEL_ID"), msgID)
			if err != nil {
				log.Fatal(err)
			}

			if len(msg.Attachments) != 1 {
				log.Fatal("Msg has more than 1 or 0 attachments")
			}
			attachment := msg.Attachments[0]

			resp, err := http.Get(attachment.URL)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fullFile = append(fullFile, data...)

			fmt.Printf("Downloaded %d/%d \n", i, len(msgIDs) - 1)
		}

		os.WriteFile(args[0], fullFile, 0644)

		return
	}


	file, err := os.ReadFile(args[0]);
	if err != nil {
		log.Fatal(err)
	}

	fileChunks := splitFile(file)

	for i, f := range fileChunks {
		msg, _ := session.ChannelMessageSendComplex(os.Getenv("CHANNEL_ID"), &discordgo.MessageSend{
			Files: []*discordgo.File{
				{
					Name: "file",
					Reader: bytes.NewReader(f),
				},
			},
		})
		addMsgID(args[0], msg.ID)
		total := len(fileChunks)
		percent := 0
		if total > 0 {
			percent = (i + 1) * 100 / total
		}
		fmt.Printf("Chunk %d/%d done. %d%% \n", i, len(fileChunks) - 1, percent)
	}
}