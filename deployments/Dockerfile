FROM --platform=$BUILDPLATFORM docker.io/golang:1.20 AS builder


# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd/context-broker

ARG TARGETARCH
ARG TARGETOS
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build


FROM registry.access.redhat.com/ubi9/ubi-minimal
WORKDIR /opt/diwise

LABEL "org.opencontainers.image.source" = "https://github.com/diwise/context-broker"

COPY --from=builder --chown=1001 /app/cmd/context-broker/context-broker /opt/diwise/
COPY --chown=1001 assets/config/authz.rego /opt/diwise/config/authz.rego
COPY --chown=1001 assets/config/default-config.yaml /opt/diwise/config/default.yaml
COPY --chown=1001 assets/jsonldcontexts/*.jsonld /opt/diwise/contexts/

RUN chown -R 1001 /opt/diwise

EXPOSE 8080
USER 1001

ENTRYPOINT ["/opt/diwise/context-broker"]
