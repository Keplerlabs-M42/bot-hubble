package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	log "github.com/gothew/l-og"
	"github.com/keplerlabsm42/hubble/internal/commands"
)

var Token string
var s *discordgo.Session

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + Token)
	if err != nil {
		log.Errorf("error creating Discord session %v", err)
		return
	}
}

var (
	commandList = []*discordgo.ApplicationCommand{
		{
			Name:        "info",
			Description: "Get info about the bot",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"info": commands.InfoCommand,
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Bot is up and running")
	})

	if err := s.Open(); err != nil {
		log.Fatalf("error opening connection %v", err)
		return
	}
	log.Info("Adding commands...\n")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commandList))
	for i, v := range commandList {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Errorf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Press Ctrl+C to exit")
	<-stop
	log.Info("Shutting down gracefully...")
}
