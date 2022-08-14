package nthclient

import "github.com/google/uuid"

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
}
