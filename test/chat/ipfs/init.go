package ipfs

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	ipfs_config "github.com/ipfs/go-ipfs-config"
	libp2p_ci "github.com/libp2p/go-libp2p-core/crypto"
	libp2p_peer "github.com/libp2p/go-libp2p-core/peer"
)

func initConfig(out io.Writer, nBitsForKeypair int) (*ipfs_config.Config, error) {
	identity, err := identityConfig(out, nBitsForKeypair)
	if err != nil {
		return nil, err
	}

	var CustomerBootstrapAddresses = []string{
		//"/ip4/47.108.166.41/tcp/4004/ws/p2p/12D3KooWKYAFaQDjyoZ1gaW6erBtgiXiTpJ9FJSEJS18cKZj4akD",
		//"/ip4/47.108.166.41/tcp/4001/p2p/12D3KooWKYAFaQDjyoZ1gaW6erBtgiXiTpJ9FJSEJS18cKZj4akD",
		//"/ip4/47.108.166.41/udp/4001/quic/p2p/12D3KooWKYAFaQDjyoZ1gaW6erBtgiXiTpJ9FJSEJS18cKZj4akD",
		"/ip4/47.108.183.230/tcp/4004/ws/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",
		"/ip4/47.108.183.230/tcp/4001/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",
		"/ip4/47.108.183.230/udp/4001/quic/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",

		"/ip4/47.108.185.233/tcp/4004/ws/p2p/12D3KooWRxvZGzeMcAbxXuomztAwn344EkmiRusF7x5H3U4RtkNN",
		"/ip4/47.108.185.233/tcp/4001/p2p/12D3KooWRxvZGzeMcAbxXuomztAwn344EkmiRusF7x5H3U4RtkNN",
		"/ip4/47.108.185.233/udp/4001/quic/p2p/12D3KooWRxvZGzeMcAbxXuomztAwn344EkmiRusF7x5H3U4RtkNN",
	}

	//bootstrapPeers, err := ipfs_config.DefaultBootstrapPeers()  // 默认启动引导项
	bootstrapPeers, err := ipfs_config.ParseBootstrapPeers(CustomerBootstrapAddresses) //定制启动引导项
	if err != nil {
		return nil, err
	}

	datastore := defaultDatastoreConfig()
	conf := &ipfs_config.Config{
		API: ipfs_config.API{
			HTTPHeaders: map[string][]string{},
		},

		// setup the node's default addresses.
		// NOTE: two swarm listen addrs, one tcp, one utp.
		Addresses: addressesConfig(),

		Datastore: datastore,
		Bootstrap: ipfs_config.BootstrapPeerStrings(bootstrapPeers),
		Identity:  identity,
		Discovery: ipfs_config.Discovery{
			MDNS: ipfs_config.MDNS{
				Enabled:  false,
				Interval: 10,
			},
		},

		Routing: ipfs_config.Routing{
			Type: "dhtclient",
		},

		// setup the node mount points.
		Mounts: ipfs_config.Mounts{
			IPFS: "/ipfs",
			IPNS: "/ipns",
		},

		Ipns: ipfs_config.Ipns{
			ResolveCacheSize: 128,
		},

		Reprovider: ipfs_config.Reprovider{
			Interval: "12h",
			Strategy: "all",
		},
		Swarm: ipfs_config.SwarmConfig{
			ConnMgr: ipfs_config.ConnMgr{
				LowWater:    defaultConnMgrLowWater,
				HighWater:   defaultConnMgrHighWater,
				GracePeriod: defaultConnMgrGracePeriod.String(),
				Type:        "basic",
			},
		},
	}

	var cc = Config{cfg: conf}
	//s := `{"Peers": [
	// {
	//   "Addrs": [
	//     "/ip4/47.108.166.41/tcp/4004",
	//     "/ip4/47.108.166.41/udp/4001"
	//   ],
	//   "ID": "12D3KooWKYAFaQDjyoZ1gaW6erBtgiXiTpJ9FJSEJS18cKZj4akD"
	// }
	//]}`
	s := `{"Peers": [
	{
        "ID": "12D3KooWRxvZGzeMcAbxXuomztAwn344EkmiRusF7x5H3U4RtkNN",
        "Addrs": [
          "/ip4/47.108.185.233/tcp/4004",
         "/ip4/47.108.185.233/tcp/4001",
	     "/ip4/47.108.185.233/udp/4001"
        ]
	},
	{
        "ID": "QmVLEz2SxoNiFnuyLpbXsH6SvjPTrHNMU88vCQZyhgBzgw",
        "Addrs": [
          "/ip4/3.215.160.238/tcp/4001",
          "/ip4/3.215.160.238/udp/4001/quic"
        ]
	},
	{
	  "Addrs": [
	     "/ip4/47.108.183.230/tcp/4001",
          "/ip4/47.108.183.230/udp/4001",
          "/ip4/47.108.183.230/tcp/4004"
	  ],
	  "ID": "12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J"
	}
	]}`
	cc.SetKey("Peering", []byte(s))
	//return conf, nil
	return cc.cfg, nil
}

// defaultConnMgrHighWater is the default value for the connection managers
// 'high water' mark
const defaultConnMgrHighWater = 200

// defaultConnMgrLowWater is the default value for the connection managers 'low
// water' mark
const defaultConnMgrLowWater = 100

// defaultConnMgrGracePeriod is the default value for the connection managers
// grace period
const defaultConnMgrGracePeriod = time.Second * 20

func addressesConfig() ipfs_config.Addresses {
	return ipfs_config.Addresses{
		Swarm: []string{
			"/ip4/0.0.0.0/tcp/0",
			"/ip6/::/tcp/0",

			"/ip4/0.0.0.0/udp/0/quic",
			"/ip6/::/udp/0/quic",
		},

		// @FIXME: use random port here to avoid collision
		// API:     ipfs_config.Strings{"/ip4/127.0.0.1/tcp/43453"},
		// Gateway: ipfs_config.Strings{"/ip4/127.0.0.1/tcp/43454"},
	}
}

// defaultDatastoreConfig is an internal function exported to aid in testing.
func defaultDatastoreConfig() ipfs_config.Datastore {
	return ipfs_config.Datastore{
		StorageMax:         "10GB",
		StorageGCWatermark: 90, // 90%
		GCPeriod:           "1h",
		BloomFilterSize:    0,
		Spec: map[string]interface{}{
			"type": "mount",
			"mounts": []interface{}{
				map[string]interface{}{
					"mountpoint": "/blocks",
					"type":       "measure",
					"prefix":     "flatfs.datastore",
					"child": map[string]interface{}{
						"type":      "flatfs",
						"path":      "blocks",
						"sync":      true,
						"shardFunc": "/repo/flatfs/shard/v1/next-to-last/2",
					},
				},
				map[string]interface{}{
					"mountpoint": "/",
					"type":       "measure",
					"prefix":     "leveldb.datastore",
					"child": map[string]interface{}{
						"type":        "levelds",
						"path":        "datastore",
						"compression": "none",
					},
				},
			},
		},
	}
}

// identityConfig initializes a new identity.
func identityConfig(out io.Writer, nbits int) (ipfs_config.Identity, error) {
	// TODO guard higher up
	ident := ipfs_config.Identity{}
	if nbits < 2048 {
		return ident, errors.New("bitsize less than 2048 is considered unsafe")
	}

	fmt.Fprintf(out, "generating %v-bit RSA keypair...", nbits)
	sk, pk, err := libp2p_ci.GenerateKeyPair(libp2p_ci.RSA, nbits)
	if err != nil {
		return ident, err
	}
	fmt.Fprintf(out, "done\n")

	// currently storing key unencrypted. in the future we need to encrypt it.
	// TODO(security)
	skbytes, err := sk.Bytes()
	if err != nil {
		return ident, err
	}
	ident.PrivKey = base64.StdEncoding.EncodeToString(skbytes)

	id, err := libp2p_peer.IDFromPublicKey(pk)
	if err != nil {
		return ident, err
	}
	ident.PeerID = id.Pretty()
	fmt.Fprintf(out, "libp2p_peer identity: %s\n", ident.PeerID)
	return ident, nil
}
