package main

import (
	"context"
	"fmt"
	"sync"

	libp2p "github.com/libp2p/go-libp2p"
	tcp "github.com/libp2p/go-tcp-transport"
	mplex "github.com/libp2p/go-libp2p-mplex"
	noise "github.com/libp2p/go-libp2p-noise"
	peer "github.com/libp2p/go-libp2p-peer"
	ma "github.com/multiformats/go-multiaddr"
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

	// Set up a listener for incoming messages
	node.SetStreamHandler("/chat/1.0.0", handleStream)

	// Parse the libp2p address of the peer we want to connect to
	peerAddr, err := ma.NewMultiaddr("/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ")
	if err != nil {
		panic(err)
	}
	peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		panic(err)
	}

	// Connect to the peer
	err = node.Connect(context.Background(), *peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to peer")

	// Wait indefinitely
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

// Handle incoming streams
func handleStream(stream libp2p.Stream) {
	// Print the incoming message
	fmt.Println(string(stream.Data()))

	// Close the stream
	stream.Close()
}
