package nthclient

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func CalculateAPIHostname(seed, tld string) string {
	t := time.Now().Truncate(0).UTC().Format("2006-01-02")
	digest := md5.Sum([]byte(seed + t))
	return fmt.Sprintf("www.%s.%s",
		hex.EncodeToString(digest[0:6]),
		tld)
}
