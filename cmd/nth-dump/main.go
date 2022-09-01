package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mdp/qrterminal/v3"

	"github.com/Snawoot/nth-dump/nthclient"
)

var version = "undefined"

var (
	// global options
	showVersion = flag.Bool("version", false, "show program version and exit")
	timeout     = flag.Duration("timeout", 30*time.Second, "operation timeout")
	format      = flag.String("format", "text", "output format: text, raw, json")
	urlFormat   = flag.String("url-format", "sip002", "output URL format: sip002, sip002u, sip002qs")
	profile     = flag.String("profile", "android", "secrets and constants profile (android/win/mac/ios)")
	loadProfile = flag.String("load-profile", "", "load JSON with settings profile from file")
	saveProfile = flag.String("save-profile", "", "save JSON profile for chosen configuration and exit")
)

func run() int {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return 0
	}

	ctx, cl := context.WithTimeout(context.Background(), *timeout)
	defer cl()

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

	if *loadProfile != "" {
		loadedSettings, err := loadSettings(*loadProfile)
		if err != nil {
			log.Fatalf("unable to load settings file: %v", err)
		}
		settings = loadedSettings
	}

	if *saveProfile != "" {
		err := saveSettings(*saveProfile, settings)
		if err != nil {
			log.Fatalf("unable to save settings file: %v", err)
		}
		return 0
	}

	nc := nthclient.New().WithSettings(settings)
	b, err := nc.GetServerConfig(ctx)
	if err != nil {
		log.Fatalf("can't get server config: %v", err)
	}

	switch *format {
	case "raw":
		fmt.Println(string(b))
	case "json":
		serverConfig, err := nthclient.UnmarshalServerConfig(b)
		if err != nil {
			log.Fatal(err)
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "    ")

		if err := enc.Encode(serverConfig.Servers); err != nil {
			log.Fatalf("can't marshal server list to json: %v", err)
		}
	default:
		serverConfig, err := nthclient.UnmarshalServerConfig(b)
		if err != nil {
			log.Fatal(err)
		}

		for _, server := range serverConfig.Servers {
			var url string
			switch *urlFormat {
			case "sip002u":
				url = server.Format(nthclient.FormatSIP002Unshielded)
			case "sip002qs":
				url = server.Format(nthclient.FormatSIP002QSAuth)
			default:
				url = server.Format(nthclient.FormatSIP002)
			}
			fmt.Println("\n----------\n")
			if !*noqr {
				qrterminal.Generate(url, qrterminal.L, os.Stdout)
			}
			fmt.Printf("Name:\t\t%s\n", server.Name)
			fmt.Printf("Host:\t\t%s\n", server.Host)
			fmt.Printf("Port:\t\t%d\n", server.Port)
			fmt.Printf("Method:\t\t%s\n", server.Method)
			fmt.Printf("Password:\t%s\n", server.Password)
			fmt.Printf("URL:\t\t%s\n", url)
		}
		fmt.Println("\n----------\n")
		if !*nowait {
			fmt.Fprintln(os.Stderr, "Press ENTER to exit...")
			fmt.Scanln()
		}
	}

	return 0
}

func loadSettings(filename string) (*nthclient.Settings, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state nthclient.Settings
	dec := json.NewDecoder(file)
	err = dec.Decode(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func saveSettings(filename string, state *nthclient.Settings) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	err = enc.Encode(state)
	return err
}

func main() {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Default().SetPrefix("NTH-DUMP: ")
	os.Exit(run())
}
