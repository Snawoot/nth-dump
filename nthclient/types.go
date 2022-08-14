package nthclient

import (
	"encoding/json"
	"fmt"
)

type ServerDefinition struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port,string"`
	Method   string `json:"method"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DomainSeedTLDDefinition struct {
	Seed string `json:"seed"`
	TLD  string `json:"tld"`
}

type ServerConfigResponse struct {
	Servers                 []*ServerDefinition        `json:"servers"`
	DomainSeed              string                     `json:"domainSeed,omitempty"`
	DomainSeedTLD           []*DomainSeedTLDDefinition `json:"domainSeedTLD,omitempty"`
	FilterdFeaturedNewsHost string                     `json:"filterdFeaturedNewsHost,omitempty"`
	OFUInterval             int64                      `json:"ofuInterval,string"`
	OFUMax                  int64                      `json:"ofuMax,string"`
}

func UnmarshalServerConfig(input []byte) (*ServerConfigResponse, error) {
	var serverConfig ServerConfigResponse

	err := json.Unmarshal(input, &serverConfig)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshalling failed: %w", err)
	}

	return &serverConfig, nil
}
