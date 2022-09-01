# nth-dump

[nthLink](https://www.nthlink.com/) API client. Retrieves shadowsocks servers and credentials. Can generate SIP-002 compatible URLs and QR codes corresponding to such URLs.

![Screenshot](https://user-images.githubusercontent.com/3524671/184556478-aaffc263-13ff-4e6f-9b3f-2dfda87cf88b.png)

## Features

* Plug and Play approach: bring your favorite shadowsocks client!
* Cross-platform (Windows/Mac OS/Linux/Android (via shell)/\*BSD)
* Zero configuration
* Simple and straightforward

## Installation

#### Binaries

Pre-built binaries are available [here](https://github.com/Snawoot/nth-dump/releases/latest).

#### Build from source

Alternatively, you may install nth-dump from source. Run the following within the source directory:

```
make install
```

#### Docker

```sh
docker run -it --rm yarmak/nth-dump
```

## Usage

Just run binary and it will output credentials.

## Synopsis

```
$ nth-dump -h
Usage of /home/user/go/bin/nth-dump:
  -format string
    	output format: text, raw, json (default "text")
  -load-profile string
    	load JSON with settings profile from file
  -noqr
    	do not print QR code with URL
  -nowait
    	do not wait for key press after output (default true)
  -profile string
    	secrets and constants profile (android/win/mac/ios) (default "android")
  -timeout duration
    	operation timeout (default 30s)
  -url-format string
    	output URL format: sip002, sip002u, sip002qs (default "sip002")
  -version
    	show program version and exit
```
