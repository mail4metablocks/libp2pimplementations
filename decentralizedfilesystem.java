import io.ipfs.api.IPFS;
import io.ipfs.multihash.Multihash;

import java.io.IOException;
import java.util.List;

public class FileSystem {
    public static void main(String[] args) throws IOException {
        // Set up IPFS with libp2p as the transport
        IPFS ipfs = new IPFS("/ip4/127.0.0.1/tcp/5001");

        // Add a file to IPFS using libp2p for transport
        Multihash hash = ipfs.add("Hello, world!".getBytes()).get(0);
        System.out.println("File added to IPFS with hash: " + hash.toBase58());

        // Search for and retrieve a file from IPFS using libp2p for transport
        byte[] file = ipfs.cat(hash);
        System.out.println("Retrieved file from IPFS: " + new String(file));
    }
}
