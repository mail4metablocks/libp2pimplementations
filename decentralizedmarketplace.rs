extern crate libp2p;
extern crate ipfs_api;

use std::io::Read;
use libp2p::{
    Transport,
    Muxer,
    Security,
};
use libp2p::tcp::TcpConfig;
use libp2p::mplex::MplexConfig;
use libp2p::noise::NoiseConfig;
use ipfs_api::IpfsClient;

fn main() {
    // Set up libp2p with the TCP and Mplex transports and the Noise encryption protocol
    let transport = TcpConfig::new().and_then(MplexConfig::new).and_then(NoiseConfig::new);
    let mut node = libp2p::build_development_node(transport).unwrap();

    // Set up IPFS with libp2p as the transport
    let mut client = IpfsClient::new(format!("/ip4/127.0.0.1/tcp/{}/p2p/{}", node.listen_addrs()[0].to_string(), node.peer_id().to_base58()));

    // Add a product listing to IPFS using libp2p for transport
    let hash = client.add("Product 1 - $100".as_bytes()).unwrap();
    println!("Product listing added to IPFS with hash: {}", hash);

    // Search for and retrieve a product listing from IPFS using libp2p for transport
    let mut listing = Vec::new();
    client.cat(&hash).unwrap().read_to_end(&mut listing).unwrap();
    println!("Retrieved product listing from IPFS: {}", String::from_utf8(listing).unwrap());
}
      
