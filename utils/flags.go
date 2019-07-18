package utils

import (
	"flag"
	"strings"
)

func ServerURL() string {
	return strings.TrimRight(flag.Lookup("micromdmurl").Value.(flag.Getter).Get().(string), "/")
}

func ApiKey() string {
	return flag.Lookup("micromdmapikey").Value.(flag.Getter).Get().(string)
}

func DebugMode() bool {
	return flag.Lookup("debug").Value.(flag.Getter).Get().(bool)
}

func Sign() bool {
	return flag.Lookup("sign").Value.(flag.Getter).Get().(bool)
}

func KeyPassword() string {
	return flag.Lookup("password").Value.(flag.Getter).Get().(string)
}

func KeyPath() string {
	return flag.Lookup("private-key").Value.(flag.Getter).Get().(string)
}

func CertPath() string {
	return flag.Lookup("cert").Value.(flag.Getter).Get().(string)
}
