# TODO: Maintain this as a seperate build container?
FROM golang:1.11 as build

# From the makefile
ARG TAGS
ARG LDFLAGS

WORKDIR /app/src/github.com/azuqua/pontus
ENV GOPATH=/app

# Copy over the source code
COPY . .

# Build the binaries
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /builddir/pontusd -a -installsuffix cgo -tags "$TAGS" -ldflags "$LDFLAGS" github.com/jaredallard/k8s-mc-helper/cmd/mchelperd

FROM debian:stable-slim
WORKDIR /srv/app

# Setup the container
RUN apt-get update && apt-get install -y ca-certificates bash \
&&  rm -rf /var/lib/apt/lists \
&&  addgroup --gid 1000 mchelper \
&&  adduser --uid 1000 --gid 1000 --shell /bin/sh --disabled-password mchelper

# Copy over the build binaries
COPY --from=build /builddir/ /usr/bin

# Setup the base system
COPY rootfs/bin/init /bin/init
RUN chmod +x /bin/init

# Run as the user pontus
USER mchelper
CMD [ "/bin/init" ]
