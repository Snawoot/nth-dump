package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Snawoot/nth-dump/nthclient"
)

var version = "undefined"

var (
	// global options
	showVersion = flag.Bool("version", false, "show program version and exit")
	timeout     = flag.Duration("timeout", 10*time.Second, "operation timeout")
)

func run() int {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return 0
	}

	ctx, cl := context.WithTimeout(context.Background(), *timeout)
	defer cl()

	nc := nthclient.New()
	_, err := nc.GetServerConfig(ctx)
	if err != nil {
		log.Fatalf("can't get server config: %v", err)
	}

	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Default().SetPrefix("NTH-DUMP: ")
	os.Exit(run())
}
