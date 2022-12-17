import io.ipfs.api.IPFS;
import io.ipfs.api.NamedStreamable;
import io.ipfs.multihash.Multihash;
import io.libp2p.core.PeerId;
import io.libp2p.core.Stream;
import io.libp2p.core.multistream.Multistream;

import java.io.IOException;
import java.util.concurrent.CompletableFuture;

public class Messaging {
    public static void main(String[] args) throws Exception {
        // Set up IPFS with libp2p as the transport
        IPFS ipfs = new IPFS("/ip4/127.0.0.1/tcp/5001");

        // Parse the libp2p address of the peer we want to connect to
        PeerId peerId = PeerId.createFromB58String("QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ");

        // Connect to the peer and set up a stream
        CompletableFuture<Stream> streamFuture = ipfs.swarm().connect(peerId);
        Stream stream = streamFuture.get();

        // Select the '/chat/1.0.0' protocol on the stream
        Multistream multistream = new Multistream(stream);
        multistream.handle("/chat/1.0.0", (protocol) -> {
            // Send a message to the peer
            protocol.write("Hello, world!".getBytes());

            // Read and print the response from the peer
            byte[] response = protocol.read();
            System.out.println(new String(response));

            return CompletableFuture.completedFuture(null);
        });
        multistream.select("/chat/1.0.0").get();

        // Close the stream
        stream.close().get();
    }
}
