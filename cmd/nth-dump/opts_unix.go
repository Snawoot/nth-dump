// +build !windows

package main

var (
	noqr        = flag.Bool("noqr", false, "do not print QR code with URL")
	nowait      = flag.Bool("nowait", true, "do not wait for key press after output")
)
