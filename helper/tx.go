package helper

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/maticnetwork/heimdall/contracts/rootchain"
	"github.com/maticnetwork/heimdall/contracts/slashmanager"

	ethereum "github.com/maticnetwork/bor"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/crypto"
	"github.com/maticnetwork/bor/ethclient"
)

func GenerateAuthObj(client *ethclient.Client, address common.Address, data []byte) (auth *bind.TransactOpts, err error) {
	// generate call msg
	callMsg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	// get priv key
	pkObject := GetPrivKey()

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.ToECDSA(pkObject[:])
	if err != nil {
		return
	}

	// from address
	fromAddress := common.BytesToAddress(pkObject.PubKey().Address().Bytes())
	// fetch gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	// fetch nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	// fetch gas limit
	callMsg.From = fromAddress
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)

	// create auth
	auth = bind.NewKeyedTransactor(ecdsaPrivateKey)
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = gasLimit

	return
}

// SendCheckpoint sends checkpoint to rootchain contract
// todo return err
func (c *ContractCaller) SendCheckpoint(signedData []byte, sigs [][3]*big.Int, rootChainAddress common.Address, rootChainInstance *rootchain.Rootchain) (er error) {
	data, err := c.RootChainABI.Pack("submitCheckpoint", signedData, sigs)
	if err != nil {
		Logger.Error("Unable to pack tx for submitCheckpoint", "error", err)
		return err
	}

	auth, err := GenerateAuthObj(GetMainClient(), rootChainAddress, data)
	if err != nil {
		Logger.Error("Unable to create auth object", "error", err)
		Logger.Info("Setting custom gaslimit", "gaslimit", GetConfig().MainchainGasLimit)
		auth.GasLimit = GetConfig().MainchainGasLimit
	}

	s := make([]string, 0)
	for i := 0; i < len(sigs); i++ {
		s = append(s, fmt.Sprintf("[%s,%s,%s]", sigs[i][0].String(), sigs[i][1].String(), sigs[i][2].String()))
	}

	Logger.Debug("Sending new checkpoint",
		"sigs", strings.Join(s, ","),
		"data", hex.EncodeToString(signedData),
	)

	tx, err := rootChainInstance.SubmitCheckpoint(auth, signedData, sigs)
	if err != nil {
		Logger.Error("Error while submitting checkpoint", "error", err)
		return err
	}
	Logger.Info("Submitted new checkpoint to rootchain successfully", "txHash", tx.Hash().String())
	return
}

// SendTick sends slash tick to rootchain contract
func (c *ContractCaller) SendTick(signedData []byte, sigs []byte, slashManagerAddress common.Address, slashManagerInstance *slashmanager.Slashmanager) (er error) {
	data, err := c.SlashManagerABI.Pack("updateSlashedAmounts", signedData, sigs)
	if err != nil {
		Logger.Error("Unable to pack tx for updateSlashedAmounts", "error", err)
		return err
	}

	auth, err := GenerateAuthObj(GetMainClient(), slashManagerAddress, data)
	if err != nil {
		Logger.Error("Unable to create auth object", "error", err)
		Logger.Info("Setting custom gaslimit", "gaslimit", GetConfig().MainchainGasLimit)
		auth.GasLimit = GetConfig().MainchainGasLimit
	}

	Logger.Info("Sending new tick",
		"sigs", hex.EncodeToString(sigs),
		"data", hex.EncodeToString(signedData),
	)

	tx, err := slashManagerInstance.UpdateSlashedAmounts(auth, signedData, sigs)
	if err != nil {
		Logger.Error("Error while submitting tick", "error", err)
		return err
	}
	Logger.Info("Submitted new tick to slashmanager successfully", "txHash", tx.Hash().String())
	return
}
