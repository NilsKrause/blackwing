package cmd

import (
	"github.com/pkg/errors"
)

var ErrAlreadyRegistered = errors.New("Each prefix has to be uique. This one already exists.")
var ErrIsSelf = errors.New("Sender of the message was the bot itself.")
var ErrIsBot = errors.New("Sender of the message was another bot.")
var ErrNotRelevant = errors.New("This message is not a command.")
var ErrMaxDepth = errors.New("Reached max depth of message words.")
