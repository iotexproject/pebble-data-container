package blockchain

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/iotexproject/iotex-blockchain-iot/util"
)

func ExecuteContract(topic string, data []byte) error {
	pwd := util.MustFetchNonEmptyParam("VAULT_PASSWORD")
	account, err := util.GetVaultAccount(pwd)
	if err != nil {
		return err
	}
	// verify the account matches the reward address
	if account.Address().String() != util.MustFetchNonEmptyParam("VAULT_ADDRESS") {
		return fmt.Errorf("key and address do not match")
	}

	endpoint := util.MustFetchNonEmptyParam("IO_ENDPOINT")
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), account)
	return executeContract(c, data)
}

func executeContract(c iotex.AuthedClient, data []byte) error {
	cstring := util.MustFetchNonEmptyParam("CONTRACT_ADDRESS")
	caddr, err := address.FromString(cstring)
	if err != nil {
		return err
	}

	// call contract to save and exec action
	ctx := context.Background()
	// TODO: Determent ABI
	abiJson, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return err
	}

	// TODO: Get some params
	gasPriceStr := util.MustFetchNonEmptyParam("GAS_PRICE")
	gasPrice, ok := big.NewInt(0).SetString(gasPriceStr, 10)
	if !ok {
		return fmt.Errorf("failed to convert string to big int")
	}
	gasLimitStr := util.MustFetchNonEmptyParam("GAS_LIMIT")
	gasLimit, err := strconv.Atoi(gasLimitStr)
	if err != nil {
		return err
	}

	// TODO: Determent contract method and params
	params := []interface{}{}
	h, err := c.Contract(caddr, abiJson).Execute("XXXMethod", params...).
		SetGasPrice(gasPrice).SetGasLimit(uint64(gasLimit)).Call(ctx)
	if err != nil {
		return err
	}
	sleepIntervalStr := util.MustFetchNonEmptyParam("SLEEP_INTERVAL")
	sleepInterval, err := strconv.Atoi(sleepIntervalStr)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(sleepInterval) * time.Second)

	resp, err := c.API().GetReceiptByAction(ctx, &iotexapi.GetReceiptByActionRequest{
		ActionHash: hex.EncodeToString(h[:]),
	})
	if err != nil {
		return err
	}
	if resp.ReceiptInfo.Receipt.Status != 1 {
		return fmt.Errorf("distributeRewards filed: %x", h)
	}
	return nil
}
