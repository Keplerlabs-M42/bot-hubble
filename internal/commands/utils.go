package commands

import "github.com/bwmarrin/discordgo"

type OptionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

func ParseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) OptionMap {
	optionMap := make(OptionMap)
	for _, option := range options {
		optionMap[option.Name] = option
	}
	return optionMap
}
