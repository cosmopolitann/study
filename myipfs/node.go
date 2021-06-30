// ready to use gomobile package for ipfs

// This package intend to only be use with gomobile bind directly if you
// want to use it in your own gomobile project, you may want to use host/node package directly

package myipfs

// Main API exposed to the ios/android

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/ipfs/go-ipfs/core/corehttp"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	sockets "github.com/libp2p/go-socket-activation"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"

	ipfs_bs "github.com/ipfs/go-ipfs/core/bootstrap"
	// ipfs_log "github.com/ipfs/go-log"
)

type Node struct {
	listeners    []manet.Listener
	muListeners  sync.Mutex
	ipfsMobile   *IpfsMobile
	friendTopics map[string]*pubsub.Topic
	// ob           *orbit
}

func NewNode(r *Repo) (*Node, error) {
	ctx := context.Background()

	if _, err := loadPlugins(r.mr.Path); err != nil {
		return nil, err
	}

	ipfscfg := &IpfsConfig{
		RepoMobile: r.mr,
		ExtraOpts: map[string]bool{
			"pubsub": true, // enable experimental pubsub feature by default
			"ipnsps": true,  // Enable IPNS record distribution through pubsub by default
		},
	}

	mnode, err := NewMobileNode(ctx, ipfscfg)
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
		// ob:           NewOrbit(ctx, mnode.IpfsNode, mnode.Repo.Path),
	}, nil
}

func (n *Node) Close() error {
	n.muListeners.Lock()
	for _, l := range n.listeners {
		l.Close()
	}
	n.muListeners.Unlock()
	// n.ob.Kv.Close()
	// n.ob.Db.Close()
	return n.ipfsMobile.Close()
}

func (n *Node) ServeUnixSocketAPI(sockpath string) (err error) {
	_, err = n.ServeMultiaddr("/unix/" + sockpath)
	return
}

// ServeTCPAPI on the given port and return the current listening maddr
func (n *Node) ServeTCPAPI(port string) (string, error) {
	return n.ServeMultiaddr("/ip4/127.0.0.1/tcp/" + port)
}

func (n *Node) ServeConfigAPI() error {
	cfg, err := n.ipfsMobile.Repo.Config()
	if err != nil {
		return err
	}

	if len(cfg.Addresses.API) > 0 {
		for _, maddr := range cfg.Addresses.API {
			if _, err := n.ServeMultiaddr(maddr); err != nil {
				log.Printf("cannot serve `%s`: %s", maddr, err.Error())
			}
		}
	}

	return nil
}

// ServeHTTPGateway collects options, creates listener, prints status message and starts serving requests
func (n *Node) ServeHTTPGateway(port string) (string, error) {
	writable := false

	listeners, err := sockets.TakeListeners("io.ipfs.gateway")
	if err != nil {
		return "", fmt.Errorf("serveHTTPGateway: socket activation failed: %s", err)
	}

	listenerAddrs := make(map[string]bool, len(listeners))
	for _, listener := range listeners {
		listenerAddrs[string(listener.Multiaddr().Bytes())] = true
	}
	gatewayAddrs := []string{"/ip4/0.0.0.0/tcp/" + port}
	for _, addr := range gatewayAddrs {
		gatewayMaddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return "", fmt.Errorf("serveHTTPGateway: invalid gateway address: %q (err: %s)", addr, err)
		}

		if listenerAddrs[string(gatewayMaddr.Bytes())] {
			continue
		}

		gwLis, err := manet.Listen(gatewayMaddr)
		if err != nil {
			return "", fmt.Errorf("serveHTTPGateway: manet.Listen(%s) failed: %s", gatewayMaddr, err)
		}
		listenerAddrs[string(gatewayMaddr.Bytes())] = true
		listeners = append(listeners, gwLis)
	}

	// we might have listened to /tcp/0 - let's see what we are listing on
	gwType := "readonly"
	if writable {
		gwType = "writable"
	}

	for _, listener := range listeners {
		fmt.Printf("Gateway (%s) server listening on %s\n", gwType, listener.Multiaddr())
	}

	//cmdctx := *cctx
	//cmdctx.Gateway = true

	var opts = []corehttp.ServeOption{
		corehttp.MetricsCollectionOption("gateway"),
		corehttp.HostnameOption(),
		corehttp.GatewayOption(writable, "/ipfs", "/ipns"),
		corehttp.VersionOption(),
		corehttp.CheckVersionOption(),
		//corehttp.CommandsROOption(cmdctx),
	}
	errc := make(chan error)
	var wg sync.WaitGroup
	for _, lis := range listeners {
		wg.Add(1)
		go func(lis manet.Listener) {
			defer wg.Done()
			errc <- corehttp.Serve(n.ipfsMobile.IpfsNode, manet.NetListener(lis), opts...)
		}(lis)
	}
	go func() {
		wg.Wait()
		close(errc)
	}()
	if len(listeners) > 0 {
		return listeners[0].Multiaddr().String(), nil
	}
	return "", fmt.Errorf("serveHTTPGateway unknown error")
}

func (n *Node) ServeMultiaddr(smaddr string) (string, error) {
	maddr, err := ma.NewMultiaddr(smaddr)
	if err != nil {
		return "", err
	}

	ml, err := manet.Listen(maddr)
	if err != nil {
		return "", err
	}

	n.muListeners.Lock()
	n.listeners = append(n.listeners, ml)
	n.muListeners.Unlock()

	go func(l net.Listener) {
		if err := n.ipfsMobile.ServeCoreHTTP(l); err != nil {
			log.Printf("serve error: %s", err.Error())
		}
	}(manet.NetListener(ml))

	return ml.Multiaddr().String(), nil
}

func init() {
	//      ipfs_log.SetDebugLogging()
}
