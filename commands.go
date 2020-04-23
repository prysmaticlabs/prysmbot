package main

var (
	// Head (current) commands.
	headSlot = &botCommand{
		group: "current",
		command: "slot",
		shorthand: "s",
		helpText: "Retrieves the current head slot of the beacon chain.",
		responseText: "Current head slot is %d.",
	}
	headEpoch = &botCommand{
		group: "current",
		command: "epoch",
		shorthand: "e",
		helpText: "Retrieves the current head epoch of the beacon chain.",
		responseText: "Current head epoch is %d.",
	}
	headJustifiedEpoch = &botCommand{
		group: "current",
		command: "justifiedEpoch",
		shorthand: "je",
		helpText: "Retrieves the current head justified epoch of the beacon chain.",
		responseText: "Current head justified epoch is %d.",
	}
	headFinalizedEpoch = &botCommand{
		group: "current",
		command: "finalizedEpoch",
		shorthand: "fe",
		helpText: "Retrieves the current head finalized epoch of the beacon chain.",
		responseText: "Current head finalized epoch is %d.",
	}
	currentValidatorCount = &botCommand{
		group: "current",
		command: "validatorCount",
		shorthand: "vc",
		helpText: "Retrieves the count of validators in the current validator set.",
		responseText: "Current validator count for epoch %d is %d.",
	}
	currentTotalBalance = &botCommand{
		group: "current",
		command: "totalBalance",
		helpText: "Retrieves the sum of all validator balances in the current validator set.",
		responseText: "Current total balance for epoch %d is %d ETH.",
	}
	currentParticipation = &botCommand{
		group: "current",
		command: "participation",
		shorthand: "p",
		helpText: "Retrieves validator participation for the last epoch.",
		responseText: "Participation for epoch %d is %.2f%%.",
	}

	// State commands.
	genesisTime = &botCommand{
		group: "state",
		command: "genesisTime",
		shorthand: "gt",
		helpText: "Retrieves genesis time from beacon state.",
		responseText: "Genesis Time is %s.",
	}
	beaconCommittee = &botCommand{
		group: "state",
		command: "committee",
		shorthand: "c",
		helpText: "Retrieves beacon committee for given slot and committee index.",
		responseText: "Committee for slot %d and index %d is %v.",
	}

	// Validator commands.
	validatorBalance = &botCommand{
		group: "val",
		command: "balance",
		shorthand: "b",
		helpText: "Retrieves balance of requested validator.",
		responseText: "Balance of validator %d is %v ETH.",
	}
	validatorActive = &botCommand{
		group: "val",
		command: "activationEpoch",
		shorthand: "ae",
		helpText: "Retrieves activation epoch of requested validator.",
		responseText: "Activation epoch of validator %d is %d.",
	}
	validatorSlashed = &botCommand{
		group: "val",
		command: "slashed",
		shorthand: "s",
		helpText: "Retrieves if the requested validator is slashed",
		responseText: "Validator index %d is %s.",
	}

	// Block commands.
	blockGraffiti = &botCommand{
		group: "block",
		command: "graffiti",
		shorthand: "g",
		helpText: "Retrieves graffiti of requested block",
		responseText: "Graffiti of block slot %d is %s.",
	}
	blockProposer = &botCommand{
		group: "block",
		command: "proposer",
		shorthand: "p",
		helpText: "Retrieves proposer index of requested block",
		responseText: "Proposer index of block slot %d is %d.",
	}
)

var (
	currentFlagGroup = &botCommandGroup{
		name: "current",
		shorthand: "c",
		displayName: "Head State Info",
		helpText: "Use %s to get info about head state.",
		flags: []*botCommand{
			headSlot,
			headEpoch,
			headJustifiedEpoch,
			headFinalizedEpoch,
			currentValidatorCount,
			currentTotalBalance,
			currentParticipation,
		},
	}
	stateFlagGroup = &botCommandGroup{
		name: "state",
		shorthand: "s",
		displayName: "Beacon State Info",
		helpText: "Use %s to query information derived from the beacon state.",
		flags: []*botCommand{
			genesisTime,
			beaconCommittee,
		},
	}
	valFlagGroup = &botCommandGroup{
		name: "val",
		shorthand: "v",
		displayName: "Validator Info",
		helpText: "Use %s to retrieve information of validators in the validator set.",
		flags: []*botCommand{
			validatorBalance,
			validatorActive,
			validatorSlashed,
		},
	}
	blockFlagGroup = &botCommandGroup{
		name: "block",
		shorthand: "b",
		displayName: "Beacon Block Info",
		helpText: "Use %s to view data on historical blocks.",
		flags: []*botCommand{
			blockGraffiti,
			blockProposer,
		},
	}
)

var allFlagGroups = []*botCommandGroup{
	currentFlagGroup,
	stateFlagGroup,
	valFlagGroup,
	blockFlagGroup,
}