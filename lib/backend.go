package qniblib

import (
	"context"
	"go.etcd.io/etcd/client"
	"log"
	"path"
	"time"
)

const (
	prefix = "/rcuda"
)
var (
	ctx = context.Background()
)

type BackEnd struct {
	etcdApi client.KeysAPI
}

func NewBackEnd(addr string) *BackEnd {
	cfg := client.Config{
		Endpoints:               []string{addr},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	kapi := client.NewKeysAPI(c)

	if err != nil {
		log.Fatal(err)
	}
	be := &BackEnd{
		etcdApi: kapi,
	}
	return be
}

func (be BackEnd) SetDevice(host, devId, state string) (err error) {
	key := path.Join(prefix, host, devId)
	log.Printf("Setting '%s' key with '%s' value", key, state)
	resp, err := be.etcdApi.Set(ctx, key, state, nil)
	if err != nil {
		return
	}
	log.Printf("Set is done. Metadata is %q\n", resp)
	return
}

func (be *BackEnd) GetDevices() (devs map[string]string, err error) {
	devs = map[string]string{}
	// fetch directory
	resp, err := be.etcdApi.Get(context.Background(), prefix, &client.GetOptions{Recursive: true, Sort: true})
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, n := range resp.Node.Nodes {
		log.Printf("Key: %q, Value: %q\n", n.Key, n.Value)
		devs[n.Key] = n.Value
	}
	return
}