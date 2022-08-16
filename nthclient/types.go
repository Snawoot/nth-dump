package nthclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	FormatSIP002 = iota
	FormatSIP002Unshielded
	FormatSIP002QSAuth
)

type ServerDefinition struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port,string"`
	Method   string `json:"method"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (sd *ServerDefinition) String() string {
	return sd.Format(FormatSIP002)
}

func (sd *ServerDefinition) Format(format int) string {
	auth := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", sd.Method, sd.Password)))
	switch format {
	case FormatSIP002Unshielded:
		return fmt.Sprintf("ss://%s:%s@%s:%d#%s", sd.Method, sd.Password, sd.Host, sd.Port, url.PathEscape(sd.Name))
	case FormatSIP002QSAuth:
		return fmt.Sprintf("ss://%s:%d?auth=%s#%s", sd.Host, sd.Port, auth, url.PathEscape(sd.Name))
	default:
		return fmt.Sprintf("ss://%s@%s:%d#%s", auth, sd.Host, sd.Port, url.PathEscape(sd.Name))
	}
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
