package main

import (
  "fmt"
  "github.com/bwmarrin/discordgo"
)

const token string = ""
var botId string
func main() {
  dg, err := discordgo.New("Bot " + token)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  u, err := dg.User("@me")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  botId = u.ID

  dg.AddHandler(messageEvent)
  err = dg.Open()
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
