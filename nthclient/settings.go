package nthclient

import (
	"crypto/rsa"

	"github.com/google/uuid"
)

const ConfigRoutePath = "/getserver-190831.php"

// Settings define important constants required for Client operation
type Settings struct {
	DomainSeed  string
	PlatformKey string
	JSONSeed    string
	TLD         string
	Language    string
	ID          string
	AppVersion  string
	UserAgent   string
	PublicKey   *rsa.PublicKey
}

// DefaultSettings is Settings with working defaults
var DefaultSettings = &Settings{
	DomainSeed:  "ewriWabKW6aMTa2W7vFNxKqgUutgpWwH",
	PlatformKey: "jk8Gh9wweC4gF8et",
	JSONSeed:    "Gu82kdDgus0248gzkqpsl948ab7a8dse",
	TLD:         "info",
	Language:    "en-US",
	ID:          uuid.Must(uuid.NewRandom()).String(),
	AppVersion:  "5.0.0",
	UserAgent:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) nthLink/5.0.0 Chrome/78.0.3905.1 Electron/7.0.0 Safari/537.36",
	PublicKey:   DefaultPublicKey(),
}
