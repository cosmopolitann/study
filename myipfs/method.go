package myipfs

import (
	"fmt"

	ipfs_core "github.com/ipfs/go-ipfs/core"
)

func GetIpfsNode(path string) (*ipfs_core.IpfsNode, error) {

	var err error
	var repo *Repo

	if !RepoIsInitialized(path) {
		config, err := NewDefaultConfig()
		if err != nil {
			fmt.Print("config error")
			return nil, err
		}
		err = InitRepo(path, config)
		if err != nil {
			fmt.Print("initRepo error")
		}

	}
	repo, err = OpenRepo(path)
	if err != nil {
		return nil, err
	}

	node, err := NewNode(repo)
	if err != nil {
		return nil, err
	}

	return node.ipfsMobile.IpfsNode, err
}
