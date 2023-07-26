FROM golang:1.21rc3-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  user

WORKDIR $GOPATH/src/grpcserver/app/

COPY . .

RUN git config --global http.sslVerify "false"

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -p 1 -o /grpcserver server.go

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /grpcserver .
COPY --from=base go/src/grpcserver/app/ .

USER user:user

EXPOSE 9001

CMD ["./grpcserver"]
