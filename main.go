package main

import (
	"github.com/NilsKrause/blackwing/cmd"
	"github.com/bwmarrin/discordgo"
)

func main() {

	var dc *discordgo.Session
	var err error

	if dc, err = discordgo.New(); err != nil {
		panic(err.Error())
	}
	defer dc.Close()

	dc.AddHandler(cmd.Handle)
}
