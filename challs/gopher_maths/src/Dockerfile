FROM golang:1.17 AS build

WORKDIR /app
COPY go.sum go.mod ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cmd

FROM scratch

COPY --from=build /app/cmd /cmd
EXPOSE 7070
ENTRYPOINT ["/cmd"]

LABEL \
    org.opencontainers.image.authors="Iver ICS" \
    org.opencontainers.image.source="https://github.com/iver-ics/iver-ctf-2021-gopher-maths" \
    org.opencontainers.image.vendor="Iver" \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.title="Iver CTF 2021 Gopher maths challenge" \
    org.opencontainers.image.description="CTF challenge using the Gopher protocol and simple arithmetics."
