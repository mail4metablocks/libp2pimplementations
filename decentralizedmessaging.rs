extern crate libp2p;
extern crate futures;

use std::error::Error;
use libp2p::{
    Transport,
    Muxer,
    Security,
};
use libp2p::tcp::TcpConfig;
use libp2p::mplex::MplexConfig;
use libp2p::noise::NoiseConfig;
use libp2p::core::upgrade;
use libp2p::core::upgrade::{Negotiated, ReadOne};
use futures::prelude::*;

fn main() -> Result<(), Box<dyn Error>> {
    // Set up libp2p with the TCP and Mplex transports and the Noise encryption protocol
    let transport = TcpConfig::new().and_then(MplexConfig::new).and_then(NoiseConfig::new);
    let mut node = libp2p::build_development_node(transport)?;

    // Set up a listener for incoming messages
    node.set_stream_handler("/chat/1.0.0", |stream| {
        // Print the incoming message
        let message = stream.read_one().wait().unwrap();
        println!("{}", String::from_utf8(message).unwrap());

        // Close the stream
        stream.close().wait().unwrap();
        Ok(())
    });

    // Parse the libp2p address of the peer we want to connect to
    let peer_addr = "/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ".parse()?;
    let peer_info = libp2p::PeerId::from_multiaddr(&peer_addr)?;

    // Connect to the peer
    let mut stream = node.new_stream().unwrap();
    stream.connect(peer_info).wait()?;

    // Send a message to the peer
    let message = b"Hello, world!";
    stream.write_all(message).wait()?;

    // Wait indefinitely
    loop {}
}
