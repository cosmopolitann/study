package example_folder

import (
	"context"
	"fmt"
	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

/// ------ Setting up the IPFS Repo

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func createTempRepo(ctx context.Context) (string, error, string) {
	//repoPath, err := ioutil.TempDir("", "ipfs-shell")
	repoPath := "./.ipfs"

	//if err != nil {
	//	return "", fmt.Errorf("failed to get temp dir: %s", err),""
	//}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err, ""
	}
	peerId := cfg.Identity.PeerID
	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err), ""
	}

	return repoPath, nil, peerId
}

/// ------ Spawning the node

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
		ExtraOpts: map[string]bool{
			"pubsub": true,
		},
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

// Spawns a node on the default repo location, if the repo exists
func spawnDefault(ctx context.Context) (icore.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		// shouldn't be possible
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err

	}

	return createNode(ctx, defaultPath)
}

// Spawns a node to be used just for this run (i.e. creates a tmp repo)
func spawnEphemeral(ctx context.Context) (icore.CoreAPI, error, string) {
	if err := setupPlugins(""); err != nil {
		return nil, err, ""
	}

	// Create a Temporary Repo
	repoPath, err, peerId := createTempRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp repo: %s", err), ""
	}

	// Spawning an ephemeral IPFS node
	api, err := createNode(ctx, repoPath)
	return api, err, peerId
}

//

func connectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peerstore.InfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peerstore.PeerInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peerstore.PeerInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()
	return nil
}

func getUnixfsFile(path string) (files.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return nil, err
	}

	f, err := files.NewReaderPathFile(path, file, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

/// -------

func InItipfs() {
	/// --- Part I: Getting a IPFS node running

	fmt.Println("-- Getting an IPFS node running -- ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using the default path (~/.ipfs), assuming that a repo exists there already
	fmt.Println("Spawning node on default repo")
	//ipfs, err := spawnDefault(ctx)
	//if err != nil {
	//	fmt.Println("No IPFS repo available on the default path")
	//}
	// Spawn a node using a temporary path, creating a temporary repo for the run
	fmt.Println("Spawning node on a temporary repo")
	ipfs, err, peerId := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}
	fmt.Println("IPFS node is running")
	fmt.Println("这是 IPFS peerId信息 ", peerId)

	/// --- Part II: Adding a file and a directory to IPFS

	fmt.Println("\n-- Adding and getting back files & directories --")

	inputBasePath := "./example-folder/"
	inputPathFile := inputBasePath + "ipfs.paper.draft3.pdf"
	inputPathDirectory := inputBasePath + "test-dir"

	someFile, err := getUnixfsNode(inputPathFile)
	if err != nil {
		panic(fmt.Errorf("Could not get File: %s", err))
	}

	cidFile, err := ipfs.Unixfs().Add(ctx, someFile)
	if err != nil {
		panic(fmt.Errorf("Could not add File: %s", err))
	}

	fmt.Printf("Added file to IPFS with CID %s\n", cidFile.String())

	someDirectory, err := getUnixfsNode(inputPathDirectory)
	if err != nil {
		panic(fmt.Errorf("Could not get File: %s", err))
	}

	cidDirectory, err := ipfs.Unixfs().Add(ctx, someDirectory)
	if err != nil {
		panic(fmt.Errorf("Could not add Directory: %s", err))
	}

	fmt.Printf("Added directory to IPFS with CID %s\n", cidDirectory.String())

	/// --- Part III: Getting the file and directory you added back

	outputBasePath := "./example-folder/"
	outputPathFile := outputBasePath + strings.Split(cidFile.String(), "/")[2]
	outputPathDirectory := outputBasePath + strings.Split(cidDirectory.String(), "/")[2]

	rootNodeFile, err := ipfs.Unixfs().Get(ctx, cidFile)
	if err != nil {
		panic(fmt.Errorf("Could not get file with CID: %s", err))
	}

	err = files.WriteTo(rootNodeFile, outputPathFile)
	if err != nil {
		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
	}

	fmt.Printf("Got file back from IPFS (IPFS path: %s) and wrote it to %s\n", cidFile.String(), outputPathFile)

	rootNodeDirectory, err := ipfs.Unixfs().Get(ctx, cidDirectory)
	if err != nil {
		panic(fmt.Errorf("Could not get file with CID: %s", err))
	}

	err = files.WriteTo(rootNodeDirectory, outputPathDirectory)
	if err != nil {
		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
	}

	fmt.Printf("Got directory back from IPFS (IPFS path: %s) and wrote it to %s\n", cidDirectory.String(), outputPathDirectory)

	/// --- Part IV: Getting a file from the IPFS Network

	fmt.Println("\n-- Going to connect to a few nodes in the Network as bootstrappers --")

	bootstrapNodes := []string{
		// IPFS Bootstrapper nodes.
		//"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		//"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		//"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		//"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		//==

		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		//"/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		//"/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		"/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWNfFx5Fgd1LjkcTT8Egz5yVT5e7orBrpGxXc8hzUKBBoA",
		"/ip4/127.0.0.1/udp/4001/quic/p2p/12D3KooWNfFx5Fgd1LjkcTT8Egz5yVT5e7orBrpGxXc8hzUKBBoA",
		"/ip4/182.150.116.150/tcp/4001/p2p/12D3KooWNfFx5Fgd1LjkcTT8Egz5yVT5e7orBrpGxXc8hzUKBBoA",


		"/ip4/47.108.183.230/tcp/4004/ws/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",
		"/ip4/47.108.183.230/tcp/4001/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",
		"/ip4/47.108.183.230/udp/4001/quic/p2p/12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J",
	}
	//peerss:=[]string{}







	go connectToPeers(ctx, ipfs, bootstrapNodes)

	exampleCIDStr := "QmUaoioqU7bxezBQZkUcgcSyokatMY71sxsALxQmRRrHrj"

	fmt.Printf("Fetching a file from the network with CID %s\n", exampleCIDStr)
	outputPath := outputBasePath + exampleCIDStr
	testCID := icorepath.New(exampleCIDStr)

	rootNode, err := ipfs.Unixfs().Get(ctx, testCID)
	if err != nil {
		panic(fmt.Errorf("Could not get file with CID: %s", err))
	}

	err = files.WriteTo(rootNode, outputPath)
	if err != nil {
		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
	}

	fmt.Printf("Wrote the file to %s\n", outputPath)

	fmt.Println("\nAll done! You just finalized your first tutorial on how to use go-ipfs as a library")

	ipfs.PubSub().Publish(ctx, "fly", []byte("于欢 喜欢  抽烟"))
	ipfs.PubSub().Publish(ctx, "fly", []byte("于欢 喜欢  打牌"))

	//
	va := `{"type":"userRegister","data":{"id":"324833623369797632","name":"22333","phone":"22333","sex":22333,"ptime":1623747170,"utime":1623747170,"nickname":"","peer_id":"22333","img":"www.baidu.com"}}`
	v1:=`{"type":"article","data":{"id":"4324","userId":"124","accesstory":"20","accesstoryType":1,"text":"1","tag":"1","playNum":3,"title":"成都","shareNum":4,"thumbnail":"刘亦菲"}}`
	v2:=`{"type":"articlePlayAdd","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","id":"408968540008222720"}}`
	//v3:=``
	v3:=`{"type":"articleShareAdd","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","id":"408968540008222720"}}`
	//

	perrs,err:=ipfs.Swarm().Peers(ctx)
	if  err!=nil{
		fmt.Println(" connet err is",err)

	}
	fmt.Println(" perrs ",perrs)


	arr,err:=ipfs.PubSub().Ls(ctx)
	if err!=nil{
		fmt.Println(" connet err is",err)

	}
	fmt.Println(" arr ",arr)

	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(time.Second)


			ipfs.PubSub().Publish(ctx, "/db-online-sync", []byte(va))
			ipfs.PubSub().Publish(ctx, "/db-online-sync", []byte(v1))
			ipfs.PubSub().Publish(ctx, "/db-online-sync", []byte(v2))
			ipfs.PubSub().Publish(ctx, "/db-online-sync", []byte(v3))
			ipfs.PubSub().Publish(ctx, "fly", []byte(v3))

		}
	}()

	//go func() {
	//	for {
	//		//监听
	//		p, err := ipfs.PubSub().Subscribe(ctx, "/db-online-sync")
	//		if err != nil {
	//			fmt.Println(" ipfs 发布 错误:", err)
	//		}
	//
	//		fmt.Println("pub:", p)
	//		p2, err := ipfs.PubSub().Subscribe(ctx, "/db-online-sync")
	//		if err != nil {
	//			fmt.Println(" ipfs 发布 错误:", err)
	//		}
	//
	//		msg1, err := p2.Next(ctx)
	//		if err != nil {
	//			fmt.Println("sub err:", err)
	//		}
	//		fmt.Println("msg1 data:", string(msg1.Data()))
	//
	//		msg, err := p.Next(ctx)
	//		if err != nil {
	//			fmt.Println("sub err:", err)
	//		}
	//		fmt.Println("msg data:", string(msg.Data()))
	//		fmt.Println("msg from peerId:", msg.From())
	//		fmt.Println("msg from Topics:", msg.Topics())
	//		fmt.Println("msg from Seq:", string(msg.Seq()))
	//		//var from =msg.From()
	//		//判断 来自 的 节点id//
	//		//msg.From()
	//
	//	}
	//	//
	//	//
	//	//		// 解析数据,调用同步方法
	//	//		// 判断 peer id 是否是自己的
	//	//		// todo
	//	//
	//	//		var sc vo.SyncParams
	//	//		err = json.Unmarshal([]byte(msg.Data()), &sc)
	//	//		if err != nil {
	//	//			sugar.Log.Error("Marshal is failed.Err is ", err)
	//	//		}
	//	//		log.Println(" 解析的 /db-online-sync  值 =", sc)
	//	//		if sc.Method == "SyncUser" {
	//	//			//json 转成 string
	//	//			jsonBytes, err := json.Marshal(sc.Data)
	//	//			if err != nil {
	//	//				fmt.Println("解析错误:", err)
	//	//			}
	//	//			fmt.Println("转换为 json 串打印结果:%s", string(jsonBytes))
	//	//			//打开数据库
	//	//			d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	//	//
	//	//			ss := sqlitedb(d)
	//	//
	//	//			resp := ss.SyncUser(string(jsonBytes))
	//	//			log.Println("这是返回的数据 =", resp)
	//	//		}
	//	//
	//	//	}
	//	//}()
	//	select {}
	//}
}
