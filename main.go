package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	log "github.com/gothew/l-og"
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
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "info",
			Description: "Get info about the bot",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"info": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey Everyone, I'm hubble bot, just here to remind you of our mortality.",
				},
			})
		},
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
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
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
