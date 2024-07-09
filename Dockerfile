ARG BASE_IMAGE=golang:alpine
ARG CONTAINER=scratch

FROM ${BASE_IMAGE}

ARG APP_NAME=whatsapp-v2

RUN apk add --update tzdata build-base
RUN addgroup -S groot && adduser -S groot -G groot # 100:101

WORKDIR /src

COPY src/go.* ./
RUN go mod download -x

COPY src ./
RUN CGO_ENABLED=1 go build -ldflags="-s -w -linkmode external -extldflags=-static" -o ${APP_NAME} -gcflags=-m -trimpath


FROM ${CONTAINER}

WORKDIR /app

COPY --from=0 etc/passwd /etc/passwd
COPY --from=0 /src/whatsapp-v2 .
COPY --from=0 /etc/ssl/certs /etc/ssl/certs
COPY --from=0 /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 3000

USER groot

ENTRYPOINT ["./whatsapp-v2"]

