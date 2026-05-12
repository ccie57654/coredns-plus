package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	_ "github.com/ccie57654/cdns-gslb-private"
	_ "github.com/coredns/rrl/plugins/rrl"
	_ "github.com/hmonsalv/coredns-doh"

	// Include all plugins.
	_ "github.com/coredns/caddy/onevent"
	_ "github.com/coredns/coredns/plugin/acl"
	_ "github.com/coredns/coredns/plugin/any"
	_ "github.com/coredns/coredns/plugin/auto"
	_ "github.com/coredns/coredns/plugin/autopath"

	//_ "github.com/coredns/coredns/plugin/azure"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/bufsize"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/cancel"
	_ "github.com/coredns/coredns/plugin/chaos"

	//_ "github.com/coredns/coredns/plugin/clouddns"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dns64"
	_ "github.com/coredns/coredns/plugin/dnssec"
	_ "github.com/coredns/coredns/plugin/dnstap"
	_ "github.com/coredns/coredns/plugin/erratic"
	_ "github.com/coredns/coredns/plugin/errors"

	//_ "github.com/coredns/coredns/plugin/etcd"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"

	//_ "github.com/coredns/coredns/plugin/geoip"
	_ "github.com/coredns/coredns/plugin/grpc"
	_ "github.com/coredns/coredns/plugin/grpc_server"
	_ "github.com/coredns/coredns/plugin/header"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/hosts"
	_ "github.com/coredns/coredns/plugin/https"
	_ "github.com/coredns/coredns/plugin/https3"

	//_ "github.com/coredns/coredns/plugin/k8s_external"
	//_ "github.com/coredns/coredns/plugin/kubernetes"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/local"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/loop"
	_ "github.com/coredns/coredns/plugin/metadata"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/minimal"
	_ "github.com/coredns/coredns/plugin/multisocket"

	//_ "github.com/coredns/coredns/plugin/nomad"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/proxyproto"
	_ "github.com/coredns/coredns/plugin/quic"
	_ "github.com/coredns/coredns/plugin/ready"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/root"

	//_ "github.com/coredns/coredns/plugin/route53"
	_ "github.com/coredns/coredns/plugin/secondary"
	_ "github.com/coredns/coredns/plugin/sign"
	_ "github.com/coredns/coredns/plugin/template"
	_ "github.com/coredns/coredns/plugin/timeouts"
	_ "github.com/coredns/coredns/plugin/tls"

	//_ "github.com/coredns/coredns/plugin/trace"
	_ "github.com/coredns/coredns/plugin/transfer"
	_ "github.com/coredns/coredns/plugin/tsig"
	_ "github.com/coredns/coredns/plugin/view"
	_ "github.com/coredns/coredns/plugin/whoami"
)

var directives = []string{
	"rrl",
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
	"autopath",
	"template",
	"hosts",
	"gslb",
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
