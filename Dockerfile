# Etapa 1: Build do Go
ARG GO_VERSION=1.20
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/ibeers/main.go

# Etapa 2: Imagem Final
FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates
        
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

USER appuser

WORKDIR /app

COPY --from=build /bin/server .
COPY --from=build /src/web ./web
COPY --from=build /src/scripts ./scripts
# Expõe a porta em que o aplicativo será executado
EXPOSE 8080

# Comando padrão para executar o aplicativo
CMD ["./server"]