package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/furrberg/chainctl/qweezxc"
	storage "github.com/furrberg/chainctl/storage"
	"github.com/spf13/cobra"
)

var (
	hash          common.Hash
	hashStr       string
	numberInt     uint64
	number        *big.Int
	url           string
	path          string
	privateKeyHex string
	addressTo     string
	tokenAmount   int
	spenderPrvK   chan ecdsa.PrivateKey
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

// deployCmd Used to deploy smart contract on Holesky Testnet.
// need to get .go file from inputted PATH and use its deploy function
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy given smart contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}

		privateKey, err := crypto.HexToECDSA(privateKeyHex)
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

		chainId, ok := new(big.Int).SetString("31337", 10)
		if !ok {
			log.Fatal("error setting chain id")
		}
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = uint64(0)  // in units
		auth.GasPrice = gasPrice

		//r := strings.NewReader(path)
		//data, err := io.ReadAll(r)
		//if err != nil {
		//	return err
		//}

		address, tx, instance, err := qweezxc.DeployQweezxc(auth, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Storage address: %v\n", address.Hex())
		fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())
		_ = instance

		return nil
	},
}

// storage smart contract interaction commands

// queryCmd Used to query fixed smart contract for its stored values.
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query for stored value in smart contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatal(err)
		}
		address := common.HexToAddress("0x8b29c91576f372F9877036c5aBc827fB41574dc8")
		instance, err := storage.NewStorage(address, client)
		if err != nil {
			log.Fatal(err)
		}

		value, err := instance.Retrieve(nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Stored value in the contract: %v\n", value)
		return nil
	},
}

// storeCmd v1: Used to store uint64(123) on the smart contract from fixed private key.
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store value in smart contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatal(err)
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
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice

		address := common.HexToAddress("0x8b29c91576f372F9877036c5aBc827fB41574dc8")
		instance, err := storage.NewStorage(address, client)
		if err != nil {
			log.Fatal(err)
		}

		valUint64 := uint64(123)

		tx, err := instance.Store(auth, new(big.Int).SetUint64(valUint64))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Tx sent: %v\n", tx.Hash().Hex())
		retrieve, err := instance.Retrieve(nil)
		if err != nil {
			return err
		}
		fmt.Printf("Retrieved value: %v\n", retrieve)
		return nil
	},
}

// erc20 token smart contract interaction commands.

// transferCmd Used to transfer specified amount of tokens from caller's wallet to given address.
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer tokens to some address",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}

		privateKey, err := crypto.HexToECDSA(privateKeyHex)
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

		chainId, ok := new(big.Int).SetString("31337", 10)
		if !ok {
			log.Fatal("error setting chain id")
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice

		address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
		instance, err := qweezxc.NewQweezxc(address, client)
		if err != nil {
			log.Fatal(err)
		}

		addressTo := common.HexToAddress(addressTo)
		tokenAmount := new(big.Int).SetInt64(int64(tokenAmount))

		balance, err := instance.BalanceOf(nil, fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		if balance.Int64() < tokenAmount.Int64() {
			log.Fatal("insufficient funds")
		}

		transfer, err := instance.Transfer(auth, addressTo, tokenAmount)
		if err != nil {
			return err
		}

		fmt.Printf("Transfer sent succesfuly!: %v\n", transfer.Hash().Hex())

		return nil
	},
}

// checkBalanceCmd Used to check the amount of ERC-20 tokens on account balance.
var checkBalanceCmd = &cobra.Command{
	Use:   "checkbalance",
	Short: "Check account balance",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}
		address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
		instance, err := qweezxc.NewQweezxc(address, client)
		if err != nil {
			log.Fatal(err)
		}

		addressTo := common.HexToAddress(addressTo)

		balance, err := instance.BalanceOf(nil, addressTo)
		if err != nil {
			return err
		}

		fmt.Printf("Current balance: %v\n", balance.Int64())

		return nil
	},
}

// approveCmd Used to set allowance of certain number of tokens for the spender by token owner
var approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve allowance for spender address",
	RunE:  approveTransfer,
}

func approveTransfer(cmd *cobra.Command, args []string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	ownerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), ownerAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, ok := new(big.Int).SetString("31337", 10)
	if !ok {
		log.Fatal("error setting chain id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := qweezxc.NewQweezxc(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	spenderAddress := common.HexToAddress(addressTo)
	amount := new(big.Int).SetInt64(int64(tokenAmount))

	tx, err := instance.Approve(auth, spenderAddress, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Owner %v\n set allowance of %v\n tokens to spender %v\n", ownerAddress, amount, spenderAddress)
	fmt.Printf("Transaction hash: %v\n", tx.Hash().Hex())

	return nil
}

// transferFromCmd Used to transfer specified amount of tokens from owner's wallet to given address.
// Called by 3rd-party address.
var transferFromCmd = &cobra.Command{
	Use:   "transfer-from",
	Short: "Transfer tokens from given to some address",
	RunE:  transferFrom,
}

func transferFrom(cmd *cobra.Command, args []string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
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
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("")
	instance, err := qweezxc.NewQweezxc(address, client)
	if err != nil {
		log.Fatal(err)
	}

	addressTo := common.HexToAddress(addressTo)
	tokenAmount := new(big.Int).SetInt64(int64(tokenAmount))

	//approve, err := instance.Approve(auth, fromAddress, tokenAmount)
	//if err != nil {
	//	return err
	//}
	//instance.Allowance(auth, fromAddress, )

	transfer, err := instance.TransferFrom(auth, fromAddress, addressTo, tokenAmount)
	if err != nil {
		return err
	}
	fmt.Printf("Transfer sent succesfuly!: %v\n", transfer.Hash().Hex())

	return nil
}

var logTCmd = &cobra.Command{
	Use:   "logT",
	Short: "Log Transfer",
	RunE:  transferLog,
}

var logACmd = &cobra.Command{
	Use:   "logA",
	Short: "Log Approve",
	RunE:  approvalLog,
}

func approvalLog(cmd *cobra.Command, args []string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	logs := make(chan types.Log, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := qweezxc.NewQweezxc(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-quit:
			return errors.New("interrupted by user")
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if err != nil {
				log.Fatal(err)
			}
			approvalInstance, err := instance.ParseApproval(vLog)
			if err != nil {
				return err
			}
			fmt.Printf("Value approved: %v\n", approvalInstance.Value)
			fmt.Printf("Owner address: %v\n", approvalInstance.Owner)
			fmt.Printf("Spender address: %v\n", approvalInstance.Spender)

		}
	}

	return nil
}

func transferLog(cmd *cobra.Command, args []string) error {

	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	logs := make(chan types.Log)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := qweezxc.NewQweezxc(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-quit:
			return errors.New("interrupted by user")
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			newLog, err := instance.ParseTransfer(vLog)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("From: %v\n", newLog.From)
			fmt.Printf("To: %v\n", newLog.To)
			fmt.Printf("Value: %v\n", newLog.Value)

		}
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "provider", "ws://127.0.0.1:8545", "provider to use")
	rootCmd.AddCommand(getTx)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(storeCmd)
	rootCmd.AddCommand(transferCmd)
	rootCmd.AddCommand(checkBalanceCmd)
	rootCmd.AddCommand(transferFromCmd)
	rootCmd.AddCommand(logTCmd)
	rootCmd.AddCommand(logACmd)
	rootCmd.AddCommand(approveCmd)
	getTx.Flags().Uint64VarP(&numberInt, "number", "n", 0, "Number of block")
	getTx.Flags().StringVar(&hashStr, "hash", "0", "Hash of transaction")
	deployCmd.Flags().StringVarP(&privateKeyHex, "privatekey", "k", "", "Private key to sign contract deployment")
	deployCmd.Flags().StringVarP(&path, "path", "p", "", "PATH to .go generated by abigen with deploy function")
	transferCmd.Flags().StringVarP(&privateKeyHex, "privatekey", "k", "", "Private key to sign a transaction")
	transferCmd.Flags().StringVarP(&addressTo, "addressto", "a", "", "Hex address to send transaction to")
	transferCmd.Flags().IntVarP(&tokenAmount, "tokens", "t", 0, "Amount of tokens to send")
	checkBalanceCmd.Flags().StringVarP(&addressTo, "addressto", "a", "", "Hex address to check balance of")
	approveCmd.Flags().StringVarP(&privateKeyHex, "privatekey", "k", "", "Private key to sign a transaction")
	approveCmd.Flags().StringVarP(&addressTo, "addressto", "a", "", "Hex address to send transaction to")
	approveCmd.Flags().IntVarP(&tokenAmount, "tokens", "t", 0, "Amount of tokens to send")
	err := deployCmd.MarkFlagRequired("privatekey")
	if err != nil {
		log.Fatal(err)
	}
	err = transferCmd.MarkFlagRequired("privatekey")
	if err != nil {
		log.Fatal(err)
	}
	err = transferCmd.MarkFlagRequired("addressto")
	if err != nil {
		log.Fatal(err)
	}
	err = transferCmd.MarkFlagRequired("tokens")
	if err != nil {
		log.Fatal(err)
	}
	err = checkBalanceCmd.MarkFlagRequired("addressto")
	if err != nil {
		log.Fatal(err)
	}
	err = approveCmd.MarkFlagRequired("addressto")
	if err != nil {
		log.Fatal(err)
	}
	err = approveCmd.MarkFlagRequired("tokens")
	if err != nil {
		log.Fatal(err)
	}
	err = approveCmd.MarkFlagRequired("privatekey")
	if err != nil {
		log.Fatal(err)
	}
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
	client, err := ethclient.Dial("https://127.0.0.1:8545")
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
