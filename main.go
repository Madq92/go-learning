package main

import (
	"fmt"
	"github.com/Madq92/etcd-demo/etcd"
)

func main() {
	s, e := etcd.GetKye("a")
	if e != nil {
		fmt.Printf("youcaa")
	}else {
		fmt.Printf("key = a ,value = %s",s)
	}

}
