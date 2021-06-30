package myipfs

import (
	"context"
	"fmt"
	"os"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
)

type orbit struct {
	Db iface.OrbitDB
	Kv orbitdb.KeyValueStore
}

func NewOrbit(ctx context.Context, ipfsnode *core.IpfsNode, path string) *orbit {
	o := new(orbit)
	path = path + "/db"
	err := Mkdir(path)
	if err != nil {
		return o
	}
	ipfs, err := coreapi.NewCoreAPI(ipfsnode)
	if err != nil {
		fmt.Println("new core api error:", err.Error())
		return o
	}
	orbit, err := orbitdb.NewOrbitDB(ctx, ipfs, &orbitdb.NewOrbitDBOptions{Directory: &path})
	if err != nil {
		fmt.Println("new orbitdb error:", err.Error())
		return nil
	}
	o.Db = orbit
	kv, err := orbit.KeyValue(ctx, "userinfo", nil)
	if err != nil {
		fmt.Println("userinfo error:", err.Error())
		return o
	}
	kv.Put(ctx, "a", []byte("liuzihua"))
	r, _ := kv.Get(ctx, "a")
	fmt.Println("成功获得数据:", string(r[:]))
	o.Kv = kv
	return o
}
func Mkdir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			fmt.Println("dir error:", err.Error())
			return err
		}
	} else {
		fmt.Println("check dir error:", err.Error())
		return err
	}
	return nil
}
