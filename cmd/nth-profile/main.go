package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/Snawoot/nth-dump/nthclient"
)

var (
	profile = flag.String("profile", "android", "secrets and constants profile (android/win/mac/ios)")
)

func run() int {
	flag.Parse()

	settings := nthclient.DefaultSettings

	switch *profile {
	case "mac":
		settings = nthclient.DefaultMacSettings
	case "win":
		settings = nthclient.DefaultWinSettings
	case "ios":
		settings = nthclient.DefaultIOSSettings
	case "android":
		settings = nthclient.DefaultAndroidSettings
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(settings)
	if err != nil {
		log.Fatalf("marshaling error: %v", err)
	}

	return 0
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Default().SetPrefix("NTH-DUMP: ")
	os.Exit(run())
}
