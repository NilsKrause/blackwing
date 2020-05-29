package char

import (
	"github.com/NilsKrause/blackwing/controller/cmd"
	"github.com/bwmarrin/discordgo"
)

var add = cmd.NewCommand()

func init() {
	add.SetExecFunc(addHandler)
}

func addHandler(s *discordgo.Session, m *discordgo.MessageCreate) error {

	return nil
}
