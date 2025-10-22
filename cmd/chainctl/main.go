package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	storage "github.com/furrberg/chainctl/build"
	"github.com/spf13/cobra"
)

var (
	hash      common.Hash
	hashStr   string
	numberInt uint64
	number    *big.Int
	url       string
)

var rootCmd = &cobra.Command{
	Use:   "chainctl",
	Short: "Chainctl",
	//PersistentPreRun: func(cmd *cobra.Command, args []string)
}

// subcommand of root to get transaction data by hash. To get this data block needs to be specified and client dialed
// by the root. if no parameter of block and client is specified, default value is used
var getTx = &cobra.Command{
	Use:   "getTx",
	Short: "Get latest transaction",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatal(err)
		}

		if numberInt == 0 {
			number = nil
		} else {
			n := new(big.Int)
			n = n.SetUint64(numberInt)
			number = n
		}

		block, err := client.BlockByNumber(context.Background(), number)
		if err != nil {
			return err
		}

		if hashStr == "0" {
			hash = block.Transactions()[0].Hash()
		} else {
			n := new(big.Int)
			n, ok := n.SetString(hashStr, 10)
			if !ok {
				return fmt.Errorf("invalid hash string: %s", hashStr)
			}
			hash = common.BigToHash(n)
		}

		tx, _, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			return err
		}

		fmt.Printf("Tx hash: %v\n", tx.Hash().Hex())

		return nil
	},
}

// deployCmd Used to deploy simple storage smart contract using private key on Holesky Testnet.
// storeCmd Used to store value via flag on the smart contract.
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy simple storage smart contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}

		privateKey, err := crypto.HexToECDSA("136d1604aebc486f4d6b71e04c66d50183b91b6d574d4213da003ef27f6652e5")
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		chainId, ok := new(big.Int).SetString("17000", 10)
		if !ok {
			log.Fatal("error setting chain id")
		}
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		address, tx, instance, err := storage.DeployStorage(auth, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Storage address: %v\n", address.Hex())
		fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
		_ = instance

		return nil
	},
}

// queryCmd Used to query fixed smart contract for its stored values.
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query for stored value in smart contract",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "provider", "https://ethereum-holesky-rpc.publicnode.com", "provider to use")
	rootCmd.AddCommand(getTx)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(queryCmd)
	getTx.Flags().Uint64VarP(&numberInt, "number", "n", 0, "Number of block")
	getTx.Flags().StringVar(&hashStr, "hash", "0", "Hash of transaction")
}

// main gets block's header, block, tx from eth network.
// - https://ethereum-holesky-rpc.publicnode.com Holesky RPC address
// - https://ethereum-rpc.publicnode.com Mainnet RPC address
func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func runE(cmd *cobra.Command, args []string) error {
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Printf("header: %v\n", header)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Printf("block: %v\n", block)

	tx, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return err
	}

	fmt.Printf("tx: %v\n", tx)

	return nil
}
