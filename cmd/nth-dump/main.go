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
	format      = flag.String("format", "text", "output format: text, raw")
	nowait      = flag.Bool("nowait", false, "do not wait for key press after output")
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
	b, err := nc.GetServerConfig(ctx)
	if err != nil {
		log.Fatalf("can't get server config: %v", err)
	}

	switch *format {
	case "raw":
		fmt.Println(string(b))
	default:
		serverConfig, err := nthclient.UnmarshalServerConfig(b)
		if err != nil {
			log.Fatal(err)
		}

		for _, server := range serverConfig.Servers {
			fmt.Println("\n----------\n")
			fmt.Printf("Name:\t\t%s\n", server.Name)
			fmt.Printf("Host:\t\t%s\n", server.Host)
			fmt.Printf("Port:\t\t%d\n", server.Port)
			fmt.Printf("Method:\t\t%s\n", server.Method)
			fmt.Printf("Password:\t%s\n", server.Password)
		}
		fmt.Println("\n----------\n")
		if !*nowait {
			fmt.Fprintln(os.Stderr, "Press ENTER to exit...")
			fmt.Scanln()
		}
	}

	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Default().SetPrefix("NTH-DUMP: ")
	os.Exit(run())
}
