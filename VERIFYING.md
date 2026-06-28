# Verifying

**NOTE the examples here only apply to v0.3.4 and newer**

## Prerequisites

**Install cosign and gh (golang assumed):**

*Debian Based*

```bash
sudo apt install gh && \
 go install github.com/sigstore/cosign/v3/cmd/cosign@latest
```

*RHEL Based*
```bash
sudo dnf -y install gh &&
 go install github.com/sigstore/cosign/v3/cmd/cosign@latest
```

You can skip `gh` prereq by just using specific versions in the below commands by using the below, replacing the Xs with your desired version:

```bash
export VERSION=vX.X.X
```

### Release Artifacts

**To verify any published release (gh auth login or GH_TOKEN populated first):**

```bash
export VERSION="$(gh release list -L 1 -R ccie57654/coredns-plus --json=tagName -q '.[] | .tagName')"
```

**Then download the `checksums.txt` and signature bundle `checksums.txt.sigstore.json` files and verify them:**

```bash
wget https://github.com/ccie57654/coredns-plus/releases/download/$VERSION/checksums.txt
wget https://github.com/ccie57654/coredns-plus/releases/download/$VERSION/checksums.txt.sigstore.json
cosign verify-blob \
    --certificate-identity "https://github.com/ccie57654/coredns-plus/.github/workflows/release.yaml@refs/tags/$VERSION" \
    --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
    --bundle "checksums.txt.sigstore.json" \
    ./checksums.txt
```

**This should succeed - you can then download any file you want from the release and verify it:**

```bash
wget "https://github.com/ccie57654/coredns-plus/releases/download/$VERSION/coredns-plus-linux-amd64.tar.gz"
sha256sum --ignore-missing -c checksums.txt
```

Output should be `OK`

**To get the SBOM, add `.sbom.json` to the end of the URL, then use `grype` to check for vulnerabilities:**

```bash
wget "https://github.com/ccie57654/coredns-plus/releases/download/$VERSION/coredns-plus-linux-amd64.tar.gz.sbom.json"
sha256sum --ignore-missing -c checksums.txt
grype sbom:coredns-plus-linux-amd64.tar.gz.sbom.json
```

**Finally, use the `gh` CLI to verify the attestations:**

```bash
gh attestation verify \
  --owner ccie57654 \
  *.tar.gz
```

### Container Images

**Signature:**

```bash
cosign verify \
  --certificate-identity "https://github.com/ccie57654/coredns-plus/.github/workflows/release.yaml@refs/tags/$VERSION" \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  "ghcr.io/ccie57654/coredns-plus:$VERSION"
```

**Check Vulnerabilities:**

```bash
grype "docker:ghcr.io/ccie57654/coredns-plus:$VERSION"
```

**Attestations**:

```bash
gh attestation verify \
  --owner ccie57654 \
  "oci://ghcr.io/ccie57654/coredns-plus:$VERSION"
```

If these checks pass you can be confident that what you are getting is what was produced by the release workflow.