package cmd

import "github.com/bwmarrin/discordgo"

type CommandFunc func(*discordgo.Session, *discordgo.MessageCreate) error

type Command struct {
	subs map[string]*Command
	exec CommandFunc

	// name of the default sub function
	defaultSub string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) GetDefault() CommandFunc {
	if ci, ok := c.subs[c.defaultSub]; ok {
		return ci.exec
	}
	return nil
}

func (c *Command) SetDefault(defaultSub string) {
	c.defaultSub = defaultSub
}

func (c *Command) AddSubCommand(prefix string, cmd *Command) error {
	if _, ok := c.subs[prefix]; !ok {
		c.subs[prefix] = cmd
		return nil
	}

	return ErrAlreadyRegistered
}

func (c *Command) SetExecFunc(exec CommandFunc) {
	c.exec = exec
}

// Handle recursivly matches the words from the message to the subcommands
// and executes the function that handles the subcommand
func (c *Command) Handle(depth int, words []string) CommandFunc {
	if depth >= len(words) {
		return nil
	}

	cmdName := words[depth]

	sub, ok := c.subs[cmdName]

	if !ok {
		// next word in the message is not a subcommand of this command,
		// execute the function of this command
		return c.exec
	}

	// if we're here that means there is a subcommand that matched the current word
	// let the subcommand handle the rest
	return sub.Handle(depth+1, words)
}
