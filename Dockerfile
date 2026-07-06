FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Two binaries: the cacher service and the one-time setup-entity tool.
RUN CGO_ENABLED=0 go build -o /out/cacher ./cmd/cacher \
 && CGO_ENABLED=0 go build -o /out/setup-entity ./cmd/setup-entity

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/cacher /usr/local/bin/cacher
COPY --from=builder /out/setup-entity /usr/local/bin/setup-entity
# setup-entity reads entities/parcel_on_chain.json relative to the workdir.
COPY entities/ ./entities/
# Default to the cacher service; run `setup-entity` by overriding the command.
CMD ["cacher"]
