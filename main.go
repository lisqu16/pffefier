package main

import (
  "fmt"
  "strings"
  "./config"
  "github.com/bwmarrin/discordgo"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "strconv"
  "time"
  "runtime"
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

  bot.UpdateStatusComplex(discordgo.UpdateStatusData{
    IdleSince: nil,
    Game: &discordgo.Game{
      Name: "jak zrobić commandhandler w go",
      Type: discordgo.GameTypeWatching,
    },
    AFK: false,
    Status: "",
  })

  connection, err := sql.Open("mysql", config.DBuser+":"+config.DBpassword+"@/"+config.DBname)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  fmt.Println("Połączono z bazą danych MySQL")
  defer connection.Close()

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

    if Cmd == config.Prefix + "info" {
      var guilds int = len(session.State.Guilds)

      embed := &discordgo.MessageEmbed{
        Author: &discordgo.MessageEmbedAuthor{},
        Color: 0x00add8,
        Title: "Statystyki 📊",
        Description: "[🔗 Repo](https://github.com/lisqu16/pffefier)",
        Fields: []*discordgo.MessageEmbedField{
          &discordgo.MessageEmbedField{
            Name: "Serwery",
            Value: strconv.Itoa(guilds),
            Inline: false,
          },
          &discordgo.MessageEmbedField{
            Name: "Wersja go",
            Value: runtime.Version(),
            Inline: true,
          },
          &discordgo.MessageEmbedField{
            Name: "Wersja discordgo",
            Value: "v0.21.0",
            Inline: true,
          },
        },
        Timestamp: time.Now().Format(time.RFC3339),
      }

      _, _ = session.ChannelMessageSendEmbed(msg.ChannelID, embed)
      return
    }
  }


}
