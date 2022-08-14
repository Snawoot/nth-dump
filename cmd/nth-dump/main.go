package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Snawoot/nth-dump/nthclient"
)

var version = "undefined"

var (
	// global options
	showVersion = flag.Bool("version", false, "show program version and exit")
)

func run() int {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return 0
	}
	log.Println(nthclient.CalculateAPIHostname(nthclient.DefaultSettings.DomainSeed, nthclient.DefaultSettings.TLD))
	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Default().SetPrefix("NTH-DUMP: ")
	os.Exit(run())
}
