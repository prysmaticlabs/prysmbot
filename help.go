package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const footerText = "Powered by the Medalla Testnet"
const footerImg = "https://prysmaticlabs.com/assets/PrysmStripe.png"

func fullHelpEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{}
	embed.Title = "PrysmBot help"
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text:    footerText,
		IconURL: footerImg,
	}

	var fields []*discordgo.MessageEmbedField
	for _, flag := range allFlagGroups {
		field := &discordgo.MessageEmbedField{
			Name: flag.displayName,
			Value: fmt.Sprintf(flag.helpText, fmt.Sprintf("`!%s.help`", flag.name)),
			Inline:  false,
		}
		fields = append(fields, field)
	}
	embed.Fields = fields
	return embed
}

func specificHelpEmbed(requestedGroup *botCommandGroup) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{}
	embed.Title = fmt.Sprintf("%s Command Help", requestedGroup.displayName)
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text:    footerText,
		IconURL: footerImg,
	}

	var fields []*discordgo.MessageEmbedField
	for _, botCommand := range requestedGroup.commands {
		field := &discordgo.MessageEmbedField{}
		if botCommand.group == randomCommandGroup.name {
			field = &discordgo.MessageEmbedField{
				Name: fmt.Sprintf("!%s", botCommand.command),
				Value: botCommand.helpText,
				Inline:  false,
			}
		} else {
			field = &discordgo.MessageEmbedField{
				Name: fmt.Sprintf("!%s.%s", requestedGroup.name, botCommand.command),
				Value: botCommand.helpText,
				Inline:  false,
			}
		}
		fields = append(fields, field)
	}
	embed.Fields = fields
	return embed
}
