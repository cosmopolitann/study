package ipfs

import (
	"fmt"

	ipfs_core "github.com/ipfs/go-ipfs/core"
)

func GetIpfsNode(path string) (*ipfs_core.IpfsNode, error) {

	repo, repoError := OpenRepo(path)
	if repoError != nil {
		return nil, repoError
	}
	if repo == nil {
		config, err := NewDefaultConfig()
		if err != nil {
			fmt.Print("config error")
			return nil, err
		}
		initRepoError := InitRepo(path, config)
		if initRepoError != nil {
			fmt.Print("initRepo error")
		}
		repo, repoError = OpenRepo(path)
		if repoError != nil {
			return nil, repoError
		}
	}
	node, nodeError := NewNode(repo)
	if nodeError != nil {
		return nil, nodeError
	}

	return node.ipfsMobile.IpfsNode, nodeError
}
