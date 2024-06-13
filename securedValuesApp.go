package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/chzyer/readline"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func connectionInitialize(
	connYaml string,
	orgName string,
	orgMSP string,
	userName string,
	channelName string,
	contractName string,
) *gateway.Contract {
	// log.Println("============ application-golang starts ============")

	// err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	// if err != nil {
	// 	log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	// }

	// walletPath := "wallet"
	// // remove any existing wallet from prior runs
	// os.RemoveAll(walletPath)
	// wallet, err := gateway.NewFileSystemWallet(walletPath)
	// if err != nil {
	// 	log.Fatalf("Failed to create wallet: %v", err)
	// }

	// if !wallet.Exists("appUser") {
	// 	err = populateWallet(wallet)
	// 	if err != nil {
	// 		log.Fatalf("Failed to populate wallet contents: %v", err)
	// 	}
	// }

	// ccpPath := filepath.Join(
	// 	"..",
	// 	"..",
	// 	"test-network",
	// 	"organizations",
	// 	"peerOrganizations",
	// 	"org1.example.com",
	// 	"connection-org1.yaml",
	// )

	// gw, err := gateway.Connect(
	// 	gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
	// 	gateway.WithIdentity(wallet, "appUser"),
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to connect to gateway: %v", err)
	// }
	// network, err := gw.GetNetwork(channelName)
	// if err != nil {
	// 	log.Fatalf("Failed to get network: %v", err)
	// }
	// contract := network.GetContract(contractName)

	// log.Println(">>> Connection and initialization successfully finished")

	// defer gw.Close()
	// return contract

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}
	if !wallet.Exists(userName) {
		err = populateWallet(wallet, orgName, orgMSP, userName)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}
	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		orgName,
		connYaml,
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, userName),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}
	contract := network.GetContract(contractName)

	log.Println(">>> Connection and initialization successfully finished")

	defer gw.Close()
	return contract
}

// func populateWallet(wallet *gateway.Wallet) error {
// 	log.Println("============ Populating wallet ============")
// 	credPath := filepath.Join(
// 		"..",
// 		"..",
// 		"test-network",
// 		"organizations",
// 		"peerOrganizations",
// 		"org1.example.com",
// 		"users",
// 		"User1@org1.example.com",
// 		"msp",
// 	)

// 	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
// 	// read the certificate pem
// 	cert, err := os.ReadFile(filepath.Clean(certPath))
// 	if err != nil {
// 		return err
// 	}

// 	keyDir := filepath.Join(credPath, "keystore")
// 	// there's a single file in this dir containing the private key
// 	files, err := os.ReadDir(keyDir)
// 	if err != nil {
// 		return err
// 	}
// 	if len(files) != 1 {
// 		return fmt.Errorf("keystore folder should have contain one file")
// 	}
// 	keyPath := filepath.Join(keyDir, files[0].Name())
// 	key, err := os.ReadFile(filepath.Clean(keyPath))
// 	if err != nil {
// 		return err
// 	}

// 	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

// 	return wallet.Put("appUser", identity)
// }

func populateWallet(wallet *gateway.Wallet, orgName string, orgMSP string, userName string) error {
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		orgName,
		"users",
		userName,
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(orgMSP, string(cert), string(key))

	return wallet.Put(userName, identity)
}

// ========================================================================

var CONTRACT map[string]*gateway.Contract

func init() {
	CONTRACT = make(map[string]*gateway.Contract)
	CONTRACT["User1@org1.example.com"] = connectionInitialize("connection-org1.yaml", "org1.example.com", "Org1MSP", "User1@org1.example.com", "mychannel", "sacc")
	CONTRACT["Admin@org1.example.com"] = connectionInitialize("connection-org1.yaml", "org1.example.com", "Org1MSP", "Admin@org1.example.com", "mychannel", "sacc")
	CONTRACT["Admin@org2.example.com"] = connectionInitialize("connection-org2.yaml", "org2.example.com", "Org2MSP", "Admin@org2.example.com", "mychannel", "sacc")
}

func ProcessPost(key, value, user string, w http.ResponseWriter) {
	if value == "" {
		log.Println("No value header was provided")
		w.Write([]byte("400 - no value header provided"))
		return
	}

	res, err := CONTRACT[user].SubmitTransaction("setSecured", key, value)

	if err != nil {
		w.Write([]byte("400 - Contract err: "))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("200 - OK -"))
	w.Write(res)
}

func ProcessGet(key, user string, w http.ResponseWriter) {
	res, err := CONTRACT[user].SubmitTransaction("getSecured", key)

	if err != nil {
		w.Write([]byte("400 - Contract err: "))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("200 - OK -"))
	w.Write(res)
}

func ProcessSecuredValue(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.Method)
	user := r.Header.Get("user")
	if user == "" {
		log.Println("No user header was provided")
		w.Write([]byte("400 - no user header provided"))
		return
	}
	log.Printf("user: %s \n", user)

	key := r.Header.Get("key")
	if user == "" {
		log.Println("No key header was provided")
		w.Write([]byte("400 - no key header provided"))
		return
	}
	log.Printf("key: %s \n", key)

	if r.Method == "GET" {
		ProcessGet(key, user, w)
	} else if r.Method == "POST" {
		ProcessPost(key, r.Header.Get("value"), user, w)
	}
}

func main() {
	http.HandleFunc("/securedValue", ProcessSecuredValue)

	err := http.ListenAndServe(":3334", nil)
	if err == nil {
		return
	}
}
