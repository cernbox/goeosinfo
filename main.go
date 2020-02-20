package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cs3org/reva/pkg/eosclient"
)

func getEosClient() *eosclient.Client {
	// we need to connect to EOS
	opts := &eosclient.Options{
		URL: *mgm,
	}

	// create eos client
	return eosclient.New(opts)
}

var (
	mgm   = flag.String("mgm", "root://eosuser-internal.cern.ch", "mgm url")
	user  = flag.String("user", "root", "user rol to execute against MGM")
	file  = flag.String("file", "/eos/user/proc", "eos path")
	inode = flag.Uint64("inode", 0, "eos inode")
	pair  = flag.Bool("pair", false, "prints inode/path pair")
)

func main() {
	flag.Parse()
	client := getEosClient()

	var info *eosclient.FileInfo
	var err error
	ctx := context.Background()
	if *inode != 0 {
		info, err = client.GetFileInfoByInode(ctx, *user, *inode)
	} else {
		info, err = client.GetFileInfoByPath(ctx, *user, *file)
	}

	if err != nil {
		fmt.Printf("%d:::ERROR\n", *inode)
		return
	}

	if *pair {
		fmt.Printf("%d:::%s\n", info.Inode, info.File)
	} else {
		fmt.Println(info)
	}

}
