package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	log "github.com/gothew/l-og"
	"github.com/joho/godotenv"
	"github.com/keplerlabsm42/hubble/internal/commands"
	"github.com/keplerlabsm42/hubble/internal/webhooks"
)

var Token string
var s *discordgo.Session

func loadEnv() {
	// Load environment variables from .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Warn("No .env file found")
	}
}

func init() {
	loadEnv()
	flag.StringVar(&Token, "t", os.Getenv("TOKEN"), "Bot token")
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

func main() {
	if err := s.Open(); err != nil {
		log.Fatalf("error opening connection %v", err)
		return
	}

	command := commands.NewCommands(s)
	command.AddHandlers()
	command.StartHandlers()
	command.RegistryCommands()
	defer s.Close()

	server := webhooks.NewServer(":3000", s)
	log.Fatal(server.Start())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Press Ctrl+C to exit")
	<-stop
	log.Info("Shutting down gracefully...")
}
