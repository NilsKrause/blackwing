package cmd

import (
	"fmt"
	"strings"

	"github.com/NilsKrause/blackwing/config"
	"github.com/bwmarrin/discordgo"
)

var baseCommand = &Command{}

// RegisterCommand is for registration of 'main' commands.
func RegisterCommand(prefix string, cmd *Command) error {
	return baseCommand.AddSubCommand(prefix, cmd)
}

// Handle is the main eventhandler for incoming messages
func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var ch *discordgo.Channel
	var err error

	if m.Author.ID == s.State.User.ID {
		fmt.Println(ErrIsSelf)
		return
	}

	if m.Author.Bot {
		fmt.Println(ErrIsBot)
		return
	}

	if ch, err = s.Channel(m.ChannelID); err != nil {
		fmt.Println(err)
		return
	}

	// if the message does not have the bot prefix
	if !strings.HasPrefix(m.Message.Content, config.Prefix()) ||
		// or is send as a dm to the bot
		ch.Type != discordgo.ChannelTypeDM {

		// we don't need to process it
		fmt.Println(ErrNotRelevant)
		return
	}

	// at this point we know the sender wants to send a command to this bot

	// split message into words to check for subcommands
	// every 'main' command is a subcommand of the bot prefix so we just split everything
	// behind the perfix into words and jam it into the base command handler
	words := strings.Split(m.Message.Content[1:], " ")
	baseCommand.Handle(0, words)(s, m)

}
