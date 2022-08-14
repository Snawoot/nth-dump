package nthclient

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const DefaultPublicKeyPEM = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA9/W/2SBYDG9rlQ3XJt39
p2ebGsZo81o1Oq6cPwP0BHIvfjeWf3l0fQaNS1zAgTenyBWNxV4Sk526mGFnnpeP
2Fjx6YMsIdULSFoz63is1Inii82DGLE5CWvzM1RvZkV8rQ5UcWRPh3je2g6Vzyd0
AKA0xxTqvQQbnsK1sEK9biMI2242yvzUEOI36M9dVr5WOzZurIC+RgE4OjAsfGNc
5rNu2ILO+T0Zq5YOiOaqh1CmvlVwlazvjUcdsEPitsMi01w4DLdAi8qJFO1dNNaE
jDFMVXT5Sxk/lmpoeRzG+aYBnd3LlIDlaaSG1ja0gxf8GHoqckLAiiV8OyDJA5Jn
ySGh0rjkuUkncmhAyrK6bEFnQYhaqxXEEUTikKhYFi0A/17JOkRXyOW/uNhS3lQo
Z42GkYlAaKSqFR4TA6nNmpup3eTyGpUKwjZqy37PT8SKytD9I1yM3No5KvtSV/lh
05yf0+JJZL0a4ChDLWa0OEuuaY/ocKO4VuVB+3KpbgfF8uAOvGBMk60QUGoG6vDK
jm2TIzxYCWojihmThx319mFytovJd/JP/c8vXVvDO4fJOYMbPhjYMju8/HmH2atE
W1dgnzDHpO9ngALzJ8XM94V0DGPvqqKg/UqOCCYZy9Zc4YofE34/7tIicI/ho4Kw
zMZ1ek4b30+kpMJ/b0xQ0UkCAwEAAQ==
-----END PUBLIC KEY-----`

func DefaultPublicKey() *rsa.PublicKey {
	block, _ := pem.Decode([]byte(DefaultPublicKeyPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}
	return pub.(*rsa.PublicKey)
}
