# Docker configuration to create Tekton Results watcher image
FROM registry.access.redhat.com/ubi8/go-toolset:1.17.7-13.1655148239 as builder

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY pkg/ pkg/
COPY cmd/ cmd/
COPY proto/ proto/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o watcher cmd/watcher/main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6-751
COPY --from=builder /opt/app-root/src/watcher /
USER 65532:65532

ENTRYPOINT ["/watcher"]
