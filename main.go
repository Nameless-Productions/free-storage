package main

import (
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


	if os.Getenv("TOKEN") == "" {
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
}