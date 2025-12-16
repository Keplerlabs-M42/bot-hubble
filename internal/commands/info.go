package commands

import "github.com/bwmarrin/discordgo"

func InfoCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey Everyone, I'm hubble bot, just here to remind you of our mortality.",
		},
	})
}
