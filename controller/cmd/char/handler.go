package char

import (
	"github.com/NilsKrause/blackwing/controller/cmd"
	"github.com/bwmarrin/discordgo"
)

var char = cmd.NewCommand()

func init() {
	char.SetExecFunc(Handle)
	char.AddSubCommand("add", add)

	if err := cmd.RegisterCommand("char", char); err != nil {
		panic(err.Error())
	}
}

// Handle is the command handler for the char command
func Handle(s *discordgo.Session, m *discordgo.MessageCreate) error {

	// this executes the default subcommand for the char command
	if defaultSub := char.GetDefault(); defaultSub != nil {
		return defaultSub(s, m)
	}

	return nil
}
