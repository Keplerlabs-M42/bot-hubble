package webhooks

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	log "github.com/gothew/l-og"
)

type Server struct {
	Port           string
	DiscordSession *discordgo.Session
}

func NewServer(port string, s *discordgo.Session) *Server {
	return &Server{
		Port:           port,
		DiscordSession: s,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/rest/webhooks/hubble", s.HandleWebhook)

	log.Infof("Starting server on port %s", s.Port)
	return http.ListenAndServe(s.Port, nil)
}
