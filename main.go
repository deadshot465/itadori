package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file.")
		return
	}
	token := os.Getenv("TOKEN")

	client, clientErr := discordgo.New("Bot " + token)
	if clientErr != nil {
		log.Fatal("Error when creating a session for Discord bot: ", clientErr)
		return
	}

	client.AddHandler(messageCreate)
	client.Identify.Intents = discordgo.IntentsGuildMessages

	clientErr = client.Open()
	if clientErr != nil {
		log.Fatal("Failed to open connection: ", clientErr)
		return
	}

	log.Printf("%s is now running. Press Ctrl-C to exit.", client.State.User.Username + "#" + client.State.User.Discriminator)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}

func messageCreate(client *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == client.State.User.ID {
		return
	}
	prefix := os.Getenv("PREFIX")

	if msg.Content == prefix + "ping" {
		client.ChannelMessageSend(msg.ChannelID, "Pong!")
	}
	
	if msg.Content == prefix + "about" {
		thumbnail := discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/attachments/700003813981028433/736189632341082152/Go-Logo_LightBlue.png",
		}

		author := discordgo.MessageEmbedAuthor{
			Name:    "Itadori Yuuji from Jujutsu Kaisen",
			IconURL: "https://cdn.discordapp.com/avatars/741107999720079471/b8830ee0ca3eed411165a99189204803.webp?size=1024",
		}
		
		footer := discordgo.MessageEmbedFooter{
			Text: "Itadori Bot: Release 0.1 | 2021-02-07",
		}

		embed := discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Description: "Itadori Yuuji in the Church of Minamoto Kou.\nItadori was inspired by the anime/manga Jujutsu Kaisen (a.k.a. Sorcery Fight).\nItadori version 0.1 was made and developed by:\n **Tetsuki Syu#1250, Kirito#9286**",
			Color:       0xD6A09A,
			Thumbnail:   &thumbnail,
			Author:      &author,
			Fields:      nil,
			Footer:      &footer,
		}

		client.ChannelMessageSendEmbed(msg.ChannelID, &embed)
	}
}