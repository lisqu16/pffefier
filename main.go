package main

import (
  "fmt"
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
  if msg.Author.ID == botId {
    return
  }

  if msg.Content == ">ping" {
    _, _ = session.ChannelMessageSend(msg.ChannelID, "Pong :ping_pong:")
    return
  }

}
