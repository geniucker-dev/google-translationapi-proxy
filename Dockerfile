FROM --platform=$BUILDPLATFORM golang:alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build,sharing=locked \
    --mount=type=cache,target=/go/pkg,sharing=locked \
    apk add --update-cache ca-certificates tzdata make && \
    go mod download && \
    \
    if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
        make translate-api args1="GOARCH=amd64"; \
    elif [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        make translate-api args1="GOARCH=arm64"; \
    elif [ "$TARGETPLATFORM" = "linux/arm/v7" ]; then \
        make translate-api args1="GOARCH=arm"; \
    else \
        echo "Unsupported platform: $TARGETPLATFORM"; \
        exit 1; \
    fi

FROM scratch

WORKDIR /app

COPY --from=builder /app/translate-api /app/translate-api
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["./translate-api"]
