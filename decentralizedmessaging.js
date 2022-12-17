const Libp2p = require('libp2p')
const TCP = require('libp2p-tcp')
const Mplex = require('libp2p-mplex')
const { NOISE } = require('libp2p-noise')

// Set up libp2p with the TCP and Mplex transports and the Noise encryption protocol
const libp2p = new Libp2p({
  transports: [new TCP()],
  connection: {
    muxer: [Mplex],
    crypto: [NOISE]
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
libp2p.handle('/chat/1.0.0', (protocol, conn) => {
  conn.on('data', (data) => {
    console.log('Received message:', data.toString())
  })
})

// Connect to a peer
libp2p.dial('/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ', '/chat/1.0.0', (err, conn) => {
  if (err) {
    throw err
  }
  console.log('Connected to peer')
  conn.write('Hello peer!')
})
