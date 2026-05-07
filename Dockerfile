FROM docker.io/library/alpine@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11 as capadd

RUN apk update && \
    apk upgrade && \
    apk add libcap

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/coredns-plus /coredns
RUN setcap cap_net_bind_service=+ep /coredns

FROM gcr.io/distroless/static:nonroot
COPY --from=capadd /coredns /coredns
USER nonroot:nonroot
WORKDIR /
EXPOSE 53 53/udp 8080 9153
ENTRYPOINT ["/coredns"]
