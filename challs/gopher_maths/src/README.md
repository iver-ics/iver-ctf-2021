# Iver's CTF 2021 challenge: Gopher maths

Relies on the Gopher protocol ([IETF RFC-1436](https://datatracker.ietf.org/doc/html/rfc1436))
where we force users to write some logic to their own Gopher client.

Users are presented with math problems. To solve them, they traverse different
Gopher menus of predefined answers.

Possible to do manually, but 1000 math problems takes a long time to do
manually.

## Config

Either configured via environment variables or via `conf.yaml` YAML file in the
current working directory.

| Env var         | YAML        | Type    | Description
| -------         | ----        | ----    | -----------
| CTF_BINDADDRESS | bindAddress | string  | Address to bind the Gopher server to and start listening from.
| CTF_FLAGFILE    | flagFile    | string  | Path to flag file containing CTF flag shown when completing the challenge
| CTF_EQUATIONS   | equations   | int     | Number of math problems to use.
| CTF_RNGSEED     | rngSeed     | int64   | Pseudo-random generator seed. Defaults to crypto-generated seed.

Look at [`config.go`](./config.go) to see the default values.

### Config YAML example

```yaml
# conf.yaml
bindAddress: "0.0.0.0:7070"
flag: "ctf{haxing-gopher-yo}"
equations: 1000
rngSeed: 415579838547
```

## Development

Requires Go v1.16 or higher.

```sh
go run .
```

Recommended to test it via a Gopher client, such as
<https://github.com/kieselsteini/delve>

## Build

### Docker build

```sh
docker build . -t iver-ctf-2021-gopher-maths

docker run --rm -it -p 7070:7070 iver-ctf-2021-gopher-maths
```
