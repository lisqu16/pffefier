package main

import (
  "fmt"
  "strings"
  "./config"
  "github.com/bwmarrin/discordgo"
)

var botId string
var bot *discordgo.Session

func main() {
  err := config.Read()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  bot, err := discordgo.New("Bot " + config.Token)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  u, err := bot.User("@me")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  botId = u.ID

  bot.AddHandler(messageEvent)
  err = bot.Open()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  var botUsername string = u.Username;

  fmt.Println(fmt.Sprintf("%d jest online!", botUsername))

  <-make(chan struct{})
  return
}

func messageEvent(session *discordgo.Session, msg *discordgo.MessageCreate) {
  if strings.HasPrefix(msg.Content, config.Prefix) {
    if msg.Author.ID == botId {
      return
    }

    Args := strings.Split(msg.Content, " ")
    Cmd := Args[0];

    if Cmd == config.Prefix + "ping" {
      _, _ = session.ChannelMessageSend(msg.ChannelID, "Pong :ping_pong:")
      return
    }
  }


}
