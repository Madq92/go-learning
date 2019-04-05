package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var client *clientv3.Client

func init() {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32775"},
		DialTimeout: 10 * time.Second,
	})
	if(err!= nil){
		log.Fatal("etcd init fail, err = ",err)
	}
	log.Fatalln("etcd init success")
}

func GetKye(key string) (string,error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := client.Get(ctx, key)
	cancelFunc()

	if err != nil {
		log.Fatal("etcd init fail, err = ",err)
		return "",err
	}

	return string(response.Kvs[0].Value),nil
}
