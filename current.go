package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gogo/protobuf/types"
	eth "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/params"
)

func getHeadCommandResult(command string) string {
	switch command {
	case headSlot.command, headSlot.shorthand:
		chainHead, err := beaconClient.GetChainHead(context.Background(), &types.Empty{})
		if err != nil {
			log.WithError(err).Error(err, "failed to get chain head")
			os.Exit(1)
		}
		return fmt.Sprintf(headSlot.responseText, chainHead.HeadSlot)
	case headEpoch.command, headEpoch.shorthand:
		chainHead, err := beaconClient.GetChainHead(context.Background(), &types.Empty{})
		if err != nil {
			log.WithError(err).Error(err, "failed to get chain head")
			os.Exit(1)
		}
		return fmt.Sprintf(headEpoch.responseText, chainHead.HeadEpoch)
	case headJustifiedEpoch.command, headJustifiedEpoch.shorthand:
		chainHead, err := beaconClient.GetChainHead(context.Background(), &types.Empty{})
		if err != nil {
			log.WithError(err).Error(err, "failed to get chain head")
			os.Exit(1)
		}
		return  fmt.Sprintf(headJustifiedEpoch.responseText, chainHead.JustifiedEpoch)
	case headFinalizedEpoch.command, headFinalizedEpoch.shorthand:
		chainHead, err := beaconClient.GetChainHead(context.Background(), &types.Empty{})
		if err != nil {
			log.WithError(err).Error(err, "failed to get chain head")
			os.Exit(1)
		}
		return fmt.Sprintf(headFinalizedEpoch.responseText, chainHead.FinalizedEpoch)
	case currentParticipation.command, currentParticipation.shorthand, currentTotalBalance.command, currentTotalBalance.shorthand:
		req := &eth.GetValidatorParticipationRequest{}
		participation, err := beaconClient.GetValidatorParticipation(context.Background(), req)
		if err != nil {
			log.WithError(err).Error(err, "failed to get chain head")
			os.Exit(1)
		}
		if command == currentParticipation.command || command == currentParticipation.shorthand{
			return fmt.Sprintf(currentParticipation.responseText, participation.Epoch, participation.Participation.GlobalParticipationRate*100)
		} else if command == currentTotalBalance.command || command == currentTotalBalance.shorthand{
			inEther := float64(participation.Participation.EligibleEther) / float64(params.BeaconConfig().GweiPerEth)
			return fmt.Sprintf(currentTotalBalance.responseText, participation.Epoch, inEther)
		}
	default:
		return ""
	}
	return ""
}
