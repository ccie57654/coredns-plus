package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	_ "github.com/coredns/rrl/plugins/rrl"
	_ "github.com/dmachard/coredns-gslb"
	_ "github.com/hmonsalv/coredns-doh"
)

var directives = []string{
	"root",
	"metadata",
	"rrl",
	"cancel",
	"tls",
	"proxyproto",
	"quic",
	"grpc_server",
	"https",
	"https3",
	"timeouts",
	"multisocket",
	"reload",
	"nsid",
	"bufsize",
	"bind",
	"debug",
	"ready",
	"health",
	"prometheus",
	"errors",
	"log",
	"dnstap",
	"local",
	"dns64",
	"any",
	"chaos",
	"tsig",
	"cache",
	"rewrite",
	"acl",
	"header",
	"dnssec",
	"gslb",
	"autopath",
	"hosts",
	"file",
	"auto",
	"secondary",
	"loop",
	"doh",
	"forward",
	"grpc",
	"erratic",
	"whoami",
	"on",
	"sign",
	"view",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
