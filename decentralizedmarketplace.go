package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	libp2p "github.com/libp2p/go-libp2p"
	tcp "github.com/libp2p/go-tcp-transport"
	mplex "github.com/libp2p/go-libp2p-mplex"
	noise "github.com/libp2p/go-libp2p-noise"
	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Set up libp2p with the TCP and Mplex transports and the Noise encryption protocol
	node, err := libp2p.New(context.Background(),
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
		libp2p.Security(noise.ID, noise.New),
	)
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// Set up IPFS with libp2p as the transport
	ipfs := ipfs.NewShell(node)

	// Add a product listing to IPFS using libp2p for transport
	hash, err := ipfs.Add(strings.NewReader("Product 1 - $100"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Product listing added to IPFS with hash:", hash)

	// Search for and retrieve a product listing from IPFS using libp2p for transport
	reader, err := ipfs.Cat(hash)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved product listing from IPFS:", string(buf))
}
