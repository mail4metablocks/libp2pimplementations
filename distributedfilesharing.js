const Libp2p = require('libp2p')
const TCP = require('libp2p-tcp')
const Mplex = require('libp2p-mplex')
const { NOISE } = require('libp2p-noise')
const IPFS = require('ipfs')

// Set up libp2p with the TCP and Mplex transports and the Noise encryption protocol
const libp2p = new Libp2p({
  transports: [new TCP()],
  connection: {
    muxer: [Mplex],
    crypto: [NOISE]
  }
})

// Set up IPFS with libp2p as the transport
const ipfs = new IPFS({
  libp2p: libp2p
})

// Start libp2p and IPFS
libp2p.start((err) => {
  if (err) {
    throw err
  }
  console.log('Libp2p started')
  ipfs.start(() => {
    console.log('IPFS started')
  })
})

// Add a file to IPFS using libp2p for transport
ipfs.files.add({
  path: 'my-file.txt',
  content: Buffer.from('Hello, world!')
}, {
  wrapWithDirectory: true,
  onlyHash: false,
  progress: (prog) => console.log(`Received: ${prog}`)
}, (err, res) => {
  if (err) {
    throw err
  }
  console.log(res)
})

// Search for and retrieve a file from IPFS using libp2p for transport
ipfs.files.get(res[0].hash, (err, files) => {
  if (err) {
    throw err
  }
  console.log(files[0].content.toString())
})
