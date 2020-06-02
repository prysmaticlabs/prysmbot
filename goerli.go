package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var depositAmount = 32
var minimumRequestAmt = depositAmount + 1
var etherInWei = big.NewInt(1000000000000000000)
var maxETHToSend = 170

var addr common.Address
var key *ecdsa.PrivateKey
var web3 *ethclient.Client

func SendGoeth(parameters []string) (string, error) {
	if len(parameters) > 2 {
		return "This command requires 2 parameters", nil
	}
	address := parameters[0]
	var amountETH int
	var err error
	if len(parameters) == 2 {
		amountETH, err = strconv.Atoi(parameters[1])
		if err != nil {
			return "", errors.Wrap(err, "could not parse")
		}
	} else {
		amountETH = maxETHToSend
	}
	if amountETH % depositAmount == 0 {
		if amountETH == depositAmount {
			amountETH = minimumRequestAmt
		} else {
			// Give an extra eth for gas fees if they request a perfect multiple.
			amountETH += 1
		}
	}
	if amountETH > maxETHToSend {
		return fmt.Sprintf("Sorry! Only up to %d GoETH at a time!", maxETHToSend), nil
	}
	if !common.IsHexAddress(address) {
		return "Please enter a valid address!", nil
	}
	toAddress := common.HexToAddress(address)

	bal, err := web3.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return "", errors.Wrap(err, "could not get account balance")
	}

	minBalance := big.NewInt(int64(amountETH))
	minBalance.Mul(minBalance, etherInWei)
	if bal.Cmp(minBalance) < 0 {
		return "Goerli Wallet is out of Ether! <@118185622543269890>", nil
	}
	nonce, err := web3.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return "", err
	}
	value := big.NewInt(0) // in wei (1 eth)
	value.Mul(etherInWei, big.NewInt(int64(amountETH)))
	gasLimit := uint64(21000) // in units
	gasPrice, err := web3.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	chainID, err := web3.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	if err != nil {
		return "", err
	}
	err = web3.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Wrap(err, "could not send")
	}
	return fmt.Sprintf("Sent! https://goerli.etherscan.io/tx/%s", signedTx.Hash().String()), nil
}

func initWallet() error {
	var err error
	if Password != "" {
		goerliKey, err := keystore.DecryptKey([]byte(EncryptedPriv), Password /*password*/)
		if err != nil {
			return err
		}
		key = goerliKey.PrivateKey
		addr = goerliKey.Address
	} else {
		key, err = crypto.HexToECDSA(EncryptedPriv)
		if err != nil {
			return err
		}
		publicKey := key.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		addr = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	web3, err = ethclient.Dial(RPCUrl)
	if err != nil {
		return err
	}

	return nil
}