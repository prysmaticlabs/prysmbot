package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	eth "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Variables used for command line parameters
var (
	Token string
	APIUrl string
	conn *grpc.ClientConn
	beaconClient eth.BeaconChainClient
	nodeClient eth.NodeClient
	log = logrus.WithField("prefix", "prysmBot")
)

func init() {
	flag.StringVar(&Token, "token", "", "Bot Token")
	flag.StringVar(&APIUrl, "api-url", "", "API Url for gRPC")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	conn, err = grpc.Dial(APIUrl, grpc.WithInsecure())
	if err != nil {
		log.Error ("Failed to dial: %v", err)
	}
	beaconClient = eth.NewBeaconChainClient(conn)
	nodeClient = eth.NewNodeClient(conn)
	defer conn.Close()

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != "691473296696410164" && m.ChannelID != "701148358445760573" {
		return
	}
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	fullCommand := m.Content[1:]
	// If the message is "ping" reply with "Pong!"
	if fullCommand == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	if fullCommand == "help" {
		embed := &discordgo.MessageEmbed{}
		embed.Title = "PrysmBot help"
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text: "Powered by the Topaz Testnet",
			IconURL: "https://prysmaticlabs.com/assets/PrysmStripe.png",
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
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			log.WithError(err).Errorf("Error sending embed %s", fullCommand)
		}
		return
	}

	if fullCommand == "food" {
		_, err := s.ChannelMessageSend(m.ChannelID, foods[rand.Int()%len(foods)])
		if err != nil {
			log.WithError(err).Errorf("Error handling command %s", fullCommand)
		}
	}

	if fullCommand == "restaurant" {
		_, err := s.ChannelMessageSend(m.ChannelID, restaurants[rand.Int()%len(foods)])
		if err != nil {
			log.WithError(err).Errorf("Error handling command %s", fullCommand)
		}
	}

	splitCommand := strings.Split(fullCommand, ".")
	if fullCommand == splitCommand[0] {
		return
	}
	commandGroup := splitCommand[0]
	endOfCommand := strings.Index(splitCommand[1], " ")
	var parameters []string
	if endOfCommand == -1 {
		endOfCommand = len(splitCommand[1])
	} else {
		parameters = strings.Split(splitCommand[1][endOfCommand:], ",")
		for i, param := range parameters {
			parameters[i] = strings.TrimSpace(param)
		}
	}
	command := splitCommand[1][:endOfCommand]

	var cmdFound bool
	var cmdGroupFound bool
	var reqGroup *botCommandGroup
	for _, flagGroup := range allFlagGroups {
		if flagGroup.name == commandGroup || flagGroup.shorthand == commandGroup {
			cmdGroupFound = true
			reqGroup = flagGroup
			for _, cmd := range reqGroup.flags {
				if command == cmd.command || command == cmd.shorthand || command == "help"{
					cmdFound = true
				}
			}
		}
	}
	if !cmdGroupFound || !cmdFound {
		return
	}

	if command == "help" {
		embed := &discordgo.MessageEmbed{}
		embed.Title = fmt.Sprintf("%s command help", reqGroup.displayName)
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text: "Powered by the Topaz Testnet",
			IconURL: "https://prysmaticlabs.com/assets/PrysmStripe.png",
		}

		var fields []*discordgo.MessageEmbedField
		for _, flag := range reqGroup.flags {
			field := &discordgo.MessageEmbedField{
				Name: fmt.Sprintf("!%s.%s", reqGroup.name, flag.command),
				Value: flag.helpText,
				Inline:  false,
			}
			fields = append(fields, field)
		}
		embed.Fields = fields
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			log.WithError(err).Errorf("Error sending embed %s", fullCommand)
			return
		}
	}

	var result string
	switch commandGroup {
	case currentFlagGroup.name, currentFlagGroup.shorthand:
		result = getHeadCommandResult(command)
	case stateFlagGroup.name, stateFlagGroup.shorthand:
		result = getStateCommandResult(command, parameters)
	case valFlagGroup.name, valFlagGroup.shorthand:
		result = getValidatorCommandResult(command, parameters)
	case blockFlagGroup.name, blockFlagGroup.shorthand:
		result = getBlockCommandResult(command, parameters)
	default:
		result = "Command not found, sorry!"
	}
	if result == "" {
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, result)
	if err != nil {
		log.WithError(err).Errorf("Error handling command %s", fullCommand)
		return
	}
}