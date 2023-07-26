# Build Stage, first pull Golang image
FROM golang:1.21rc3-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  user

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/grpcserver/app/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN git config --global http.sslVerify "false"

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -p 1 -o /grpcserver server.go

# Run Stage
FROM scratch
#FROM docker.artifactory.corp.plusnet/alpine:3.17.0
#RUN apk --update --no-cache add bash && rm -rf /var/cache/apk/*

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY ca-certs/polkomtel/ca /etc/ssl/certs
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /grpcserver .
COPY --from=base go/src/grpcserver/app/ .

USER user:user

# This container exposes port to the outside world
EXPOSE 9001

# Run the binary program produced by `go install`
CMD ["./grpcserver"]
