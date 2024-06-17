FROM golang:1.22.4-bookworm AS builder

WORKDIR /src/app
COPY . /src/app

ENV CGO_ENABLED=0
ENV CGO_LDFLAGS="-s -w"
ENV GOOS=linux

RUN go build -a -installsuffix cgo -o /bin/app cmd/connected_roots/main.go

FROM gcr.io/distroless/base-debian12:latest
COPY --from=builder /bin/app /bin/app

COPY configs/* /data/config/

CMD ["/bin/app"]
