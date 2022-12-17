const Libp2p = require('libp2p')
const TCP = require('libp2p-tcp')
const Mplex = require('libp2p-mplex')
const { NOISE } = require('libp2p-noise')
const DHT = require('libp2p-kad-dht')

// Set up libp2p with the TCP and Mplex transports, the Noise encryption protocol, and the DHT module
const libp2p = new Libp2p({
  transports: [new TCP()],
  connection: {
    muxer: [Mplex],
    crypto: [NOISE]
  },
  modules: {
    dht: new DHT()
  }
})

// Start libp2p
libp2p.start((err) => {
  if (err) {
    throw err
  }
  console.log('Libp2p started')
})

// Set up a listener for incoming messages
libp2p.handle('/social/1.0.0', (protocol, conn) => {
  conn.on('data', (data) => {
    console.log('Received message:', data.toString())
  })
})

// Find and connect to other users on the social network using the DHT
libp2p.dht.findPeer('alice', (err, peerInfo) => {
  if (err) {
    throw err
  }
  libp2p.dial(peerInfo, '/social/1.0.0', (err, conn) => {
    if (err) {
      throw err
    }
    console.log('Connected to Alice')
    conn.write('Hello Alice!')
  })
})

libp2p.dht.findPeer('bob', (err, peerInfo) => {
  if (err) {
    throw err
  }
  libp2p.dial(peerInfo, '/social/1.0.0', (err, conn) => {
    if (err) {
      throw err
    }
    console.log('Connected to Bob')
    conn.write('Hello Bob!')
  })
})
