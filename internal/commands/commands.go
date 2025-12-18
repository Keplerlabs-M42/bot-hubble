package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/gothew/l-og"
)

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Command struct {
	Session  *discordgo.Session
	Handlers map[string]CommandHandler
}

var commandRegistry = []*discordgo.ApplicationCommand{
	{
		Name:        "info",
		Description: "Get info about the bot",
	},
}

func NewCommands(session *discordgo.Session) *Command {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Bot is up and running")
	})
	return &Command{
		Session: session,
	}
}

func (c *Command) StartHandlers() {
	c.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := c.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (c *Command) AddHandlers() {
	c.Handlers = map[string]CommandHandler{
		"info": InfoCommand,
	}
}

func (c *Command) RegistryCommands() ([]*discordgo.ApplicationCommand, error) {
	log.Info("Adding commands...")
	registerCommands := make([]*discordgo.ApplicationCommand, len(commandRegistry))
	for i, v := range commandRegistry {
		cmd, err := c.Session.ApplicationCommandCreate(c.Session.State.User.ID, "", v)
		if err != nil {
			return nil, err
		}
		registerCommands[i] = cmd
	}
	return commandRegistry, nil
}
