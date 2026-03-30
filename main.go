package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	_ "github.com/coredns/coredns/core/plugin"
	"github.com/coredns/coredns/coremain"

	_ "github.com/coredns/rrl/plugins/rrl"
	_ "github.com/dmachard/coredns-gslb"
	_ "github.com/hmonsalv/coredns-doh"
)

var directives = []string{
	"root",
	"metadata",
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
	"trace",
	"ready",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"dnstap",
	"local",
	"dns64",
	"rrl",
	"acl",
	"any",
	"chaos",
	"loadbalance",
	"tsig",
	"cache",
	"rewrite",
	"header",
	"dnssec",
	"autopath",
	"minimal",
	"template",
	"transfer",
	"hosts",
	"kubernetes",
	"file",
	"gslb",
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
	"nomad",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
