FROM golang:1.22.5-bookworm AS builder

WORKDIR /src/app
COPY . /src/app

ENV CGO_ENABLED=0
ENV CGO_LDFLAGS="-s -w"
ENV GOOS=linux

RUN go build -a -installsuffix cgo -o /bin/app cmd/connected_roots/main.go

# Copy config.
COPY configs/local.yaml /data/config/local.yaml
# Copy assets.
COPY internal/connected_roots/frontend/web/assets/. /data/web/assets/
# Copy templates views.
COPY internal/connected_roots/frontend/web/views/. /data/web/views/
# Copy i18n files.
COPY internal/connected_roots/frontend/i18n/*.toml /data/web/i18n/

CMD ["/bin/app"]
