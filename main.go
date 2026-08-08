package main

import (
	"bytes"
	"fmt"
	"log"
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
		fmt.Printf("Chunk %d/%d done. %d%% \n", i, len(fileChunks) - 1, i * 100 / (len(fileChunks) - 1))
	}
}