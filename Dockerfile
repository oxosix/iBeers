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

# Use a imagem base do Golang
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
COPY --from=build /src/ascii.txt .
# Define o diretório de trabalho dentro do contêiner


ARG DB_PASSWORD=beers123456
ARG DB_PORT=5432
ARG DB_USER=beers_user
ARG DB_HOST=172.17.0.1

ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_HOST=${DB_HOST}

# Copia o código fonte do aplicativo para o diretório de trabalho no contêiner

# Compila o aplicativo
# RUN go build -o app ./cmd/ibeers

# Expõe a porta em que o aplicativo será executado
EXPOSE 8080

# Comando padrão para executar o aplicativo
CMD ["./server"]
