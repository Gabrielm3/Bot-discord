package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/gabrielm3/bot-discord/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot iniciado...")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "java" ||
		m.Content == "node" ||
		m.Content == "python" ||
		m.Content == "c++" ||
		m.Content == "c#" ||
		m.Content == "c" ||
		m.Content == "javascript" ||
		m.Content == "matrix" {

		_, _ = s.ChannelMessageSend(m.ChannelID, "Saía da Matrix, Neo!\nGo é a resposta que tem procurado há tanto tempo!")
	}
}
