package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/keplerlabsm42/hubble/pkg/types"
)

func JiraCommand(s *discordgo.Session, i *discordgo.InteractionCreate, payload *types.WebhookPayload, chId string) {

	description := "updated to status: " + payload.Issue.Fields.Status.Name + " by " + payload.User.DisplayName

	title := fmt.Sprintf("Jira Task (%s): %s", payload.Issue.Key, payload.Issue.Fields.Summary)
	s.ChannelMessageSendEmbed(chId, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Now", Value: payload.Issue.Fields.Status.Name, Inline: true},
		},
		Color: 0x00ff00,
	})
}
