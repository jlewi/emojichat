# Build the manager binary
#
# TODO(jlewi): Maybe we should build a builder image with the dependencies baked in
# to speedup rebuilds.
FROM golang:1.13 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server cmd/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /workspace/server .
COPY static /static/

ENTRYPOINT ["/server", "serve"]