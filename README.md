# CoreDNS-plus

## What is coredns-plus?

coredns-plus is just a project to test compiling **[CoreDNS](https://coredns.io/)** with additional plugins compiled in as well as fully removing some plugins that i don't have a use-case for. 

I also use it for validation of pipeline publishing and performance testing some private go modules that are just duplicates of public forks of other plugins like: **[CoreDNS-GSLB](https://github.com/dmachard/coredns-gslb/)**

Some of the specifics:
- You can see all of the included and excluded plugins in main.go which are just copied from the coredns repo
- They are explicitly declared imports because any other method of masking them out doesn't reduce the footprint and I prefer to keep the dependency tree pruned when possible
- I also jump to the latest golang whenever that comes out if there are higher sev CVEs involved

To be added soon:
- automated validation of included plugins
- testing DoH which might be getting included in the core binary soon so backburnered
- removing setcap from requirements, working out how to make sure recursors can still reach the container on :53 when there is no proxy translating
- some dnstap VRL examples for vector
- adding TLS/mTLS to grpc backends for CoreDNS-GSLB

## signing and attestation

- install cosign

```bash
go install github.com/sigstore/cosign/v3/cmd/cosign@latest
```

either sign or verify as shown [here](https://docs.sigstore.dev/quickstart/quickstart-cosign/)