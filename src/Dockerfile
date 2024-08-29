FROM golang:1.23.0-bookworm as builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app
COPY ./ ./
RUN go mod download && go build -o /app/out/gh-pkg-cli /app/main.go


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder --chown=nobody:nogroup /app/out/gh-pkg-cli /app/

ENTRYPOINT ["/app/gh-pkg-cli"]
CMD ["--help"]
