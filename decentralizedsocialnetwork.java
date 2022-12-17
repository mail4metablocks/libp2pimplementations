import io.ipfs.api.IPFS;
import io.ipfs.api.NamedStreamable;
import io.ipfs.multihash.Multihash;

import java.io.IOException;
import java.util.List;

public class SocialNetwork {
    public static void main(String[] args) throws IOException {
        // Set up IPFS with libp2p as the transport
        IPFS ipfs = new IPFS("/ip4/127.0.0.1/tcp/5001");

        // Add a post to IPFS using libp2p for transport
        Multihash hash = ipfs.add(new NamedStreamable.ByteArrayWrapper("Hello, world!".getBytes())).get(0);
        System.out.println("Post added to IPFS with hash: " + hash.toBase58());

        // Search for and retrieve a post from IPFS using libp2p for transport
        byte[] post = ipfs.cat(hash);
        System.out.println("Retrieved post from IPFS: " + new String(post));
    }
}
