FROM golang:1.17 AS build

WORKDIR /app
COPY go.sum go.mod ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cmd

FROM scratch

COPY --from=build /app/cmd /cmd
EXPOSE 8080
ENTRYPOINT ["/cmd"]

LABEL \
    org.opencontainers.image.authors="Iver ICS" \
    org.opencontainers.image.vendor="Iver"
