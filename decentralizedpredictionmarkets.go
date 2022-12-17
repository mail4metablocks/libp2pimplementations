package main

import (
	"fmt"
	"io/ioutil"

	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Set up IPFS with libp2p as the transport
	client := ipfs.NewShell("127.0.0.1:5001")

	// Add a prediction to IPFS using libp2p for transport
	hash, err := client.Add([]byte("Prediction: The sun will rise tomorrow."))
	if err != nil {
		panic(err)
	}
	fmt.Println("Prediction added to IPFS with hash:", hash)

	// Search for and retrieve a prediction from IPFS using libp2p for transport
	prediction, err := client.Cat(hash)
	if err != nil {
		panic(err)
	}
	predictionStr, err := ioutil.ReadAll(prediction)
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved prediction from IPFS:", string(predictionStr))
}
