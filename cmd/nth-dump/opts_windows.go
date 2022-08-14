// +build windows

package main

var (
	noqr        = flag.Bool("noqr", true, "do not print QR code with URL")
	nowait      = flag.Bool("nowait", false, "do not wait for key press after output")
)
