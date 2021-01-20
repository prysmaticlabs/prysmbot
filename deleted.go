package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func handleMessageDeleted(s *discordgo.Session, m *discordgo.MessageDelete) {
	if !goerliOkayChannel(m.ChannelID) {
		return
	}

	if m.BeforeDelete == nil {
		log.WithField("beforeDelete", m.BeforeDelete).Warn("Deleted message not tracked in state, not handling")
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("User %s (%s) just deleted a message: \"%s\"", m.BeforeDelete.Author.Username, m.BeforeDelete.Author.ID, m.BeforeDelete.Content))
}
