package webhooks

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/keplerlabsm42/hubble/internal/commands"
	"github.com/keplerlabsm42/hubble/pkg/types"
)

func (s *Server) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var payload types.WebhookPayload

	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// TODO: update to use channel from config
	commands.JiraCommand(s.DiscordSession, nil, &payload, "1448878479843262545")

	w.WriteHeader(http.StatusOK)
}
