FROM golang:1.13.4-alpine3.10 as builder

ENV GOPATH /workspace
ENV PATH "$PATH:/usr/local/go/bin:$GOPATH/bin"
ENV CGO_ENABLED 0
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOPATH/pkg"

ENV SRC_DIR "$GOPATH/src/org.junyan/capture"
WORKDIR $SRC_DIR
COPY *.go $SRC_DIR
COPY go.mod $SRC_DIR

RUN go build

# ============================================

FROM alpine:latest
COPY --from=builder /workspace/src/org.junyan/capture/capture /usr/local/bin/capture/capture
WORKDIR /usr/local/bin/capture
ENTRYPOINT ["/usr/local/bin/capture/capture", "--port=3000"]