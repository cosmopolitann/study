package ipfs

// Main API exposed to the ios/android

import (
	"context"
	"log"
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	ipfs_mobile "github.com/ipfs-shipyard/gomobile-ipfs/go/pkg/ipfsmobile"

	manet "github.com/multiformats/go-multiaddr/net"

	ipfs_bs "github.com/ipfs/go-ipfs/core/bootstrap"
	// ipfs_log "github.com/ipfs/go-log"
)

type Node struct {
	listeners    []manet.Listener
	muListeners  sync.Mutex
	ipfsMobile   *ipfs_mobile.IpfsMobile
	friendTopics map[string]*pubsub.Topic
}

func NewNode(r *Repo) (*Node, error) {
	ctx := context.Background()

	if _, err := loadPlugins(r.mr.Path); err != nil {
		return nil, err
	}

	ipfscfg := &ipfs_mobile.IpfsConfig{
		RepoMobile: r.mr,
		ExtraOpts: map[string]bool{
			"pubsub": false, // enable experimental pubsub feature by default
			"ipnsps": true,  // Enable IPNS record distribution through pubsub by default
		},
	}

	mnode, err := ipfs_mobile.NewNode(ctx, ipfscfg)
	if err != nil {
		return nil, err
	}

	if err := mnode.IpfsNode.Bootstrap(ipfs_bs.DefaultBootstrapConfig); err != nil {
		log.Printf("failed to bootstrap node: `%s`", err)
	}
	topics := make(map[string]*pubsub.Topic)
	return &Node{
		ipfsMobile:   mnode,
		friendTopics: topics,
	}, nil
}

func (n *Node) Close() error {
	n.muListeners.Lock()
	for _, l := range n.listeners {
		l.Close()
	}
	n.muListeners.Unlock()
	return n.ipfsMobile.Close()
}
