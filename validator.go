package main

import (
	"context"
	"fmt"
	"strconv"

	eth "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/params"
)

func getValidatorCommandResult(command string, parameters []string) string {
	if len(parameters) != 1 {
		log.Error("Expected 1 parameter for validator command")
		return ""
	}
	reqIndex, err := strconv.Atoi(parameters[0])
	if err != nil {
		log.WithError(err).Error(err, "failed to convert")
		return ""
	}
	req := &eth.GetValidatorRequest{
		QueryFilter: &eth.GetValidatorRequest_Index{
			Index: uint64(reqIndex),
		},
	}
	validator, err := beaconClient.GetValidator(context.Background(), req)
	if err != nil {
		log.WithError(err).Error(err, "failed to get committees")
		return ""
	}
	switch command {
	case validatorBalance.command, validatorBalance.shorthand:
		inEther := validator.EffectiveBalance / params.BeaconConfig().GweiPerEth
		return fmt.Sprintf(validatorBalance.responseText, reqIndex, inEther)
	case validatorActive.command, validatorActive.shorthand:
		return fmt.Sprintf(validatorActive.responseText, reqIndex, validator.ActivationEpoch)
	case validatorSlashed.command, validatorSlashed.shorthand:
		resultText := "not slashed"
		if validator.Slashed {
			resultText = "slashed"
		}
		return fmt.Sprintf(validatorSlashed.responseText, reqIndex, resultText)
	default:
		return ""
	}
}
