# X-Mas-For

- Category: `web`
- Challenge author: **Kalle Fagerberg**

## Description

Santa's senior security advisors has coded up their new credentials store, and have ensured that only Santas local networks can access it for obvious reasons. They have worked really hard on this, so you should probably give up already.

### Connection info

```sh
curl 2021.santahack.xyz 42004
```

## Building prerequisites

Go v1.16+

## Writeup

Easiest completed via HTTP utility such as the [cURL](https://curl.se/) CLI or
the [Postman](https://www.postman.com/) GUI.

A basic request returns the "403 Forbidden" HTTP status code:

```console
$ curl http://2021.santahack.xyz:42004
<!DOCTYPE html>
<html>
  <head>
    <title>Forbidden</title>
  </head>
  <body>
    <h1>Forbidden</h1>
    <p>Your IP is not in the allowed list.</p>
    <p>Santa's evil ICS elves has locked this site down to only be accessible via:</p>
    <ul>
      <li><pre>192.168.0.0/16</pre></li>
      <li><pre>127.0.0.0/8</pre></li>
    </ul>
  </body>
</html>
```

To solve it, you have to trick the application that you come from one of the
above networks (as defined by the [CIDR masks](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
in the HTML response).

Solution is to set the `X-Forwarded-For` header with an IP address within one
of the mask's ranges, such as `127.0.0.1`, like so:

```console
$ curl http://2021.santahack.xyz:42004 -H "X-Forwarded-For: 127.0.0.1"
<!DOCTYPE html>
<html>
  <head>
    <title>x_mas_for</title>
  </head>
  <body>
    <h1>x_mas_for</h1>
    <p>Welcome to our super secret web page!</p>
    <pre>haxmas{begin-ahead-of-the-headers}</pre>
  </body>
</html>
```

The reason this works is because the application lives behind a reverse proxy,
so it will always get the same IP address on it's connections (the IP address
of the reverse proxy itself), so the application needs to figure out the
end-user's IP in a different way. Different web frameworks handle this
differently, but the common solution is to take the first IP address from the
`X-Forwarded-For` header.

Nginx, and many other reverse proxies, append the connecting IP address to the
`X-Forwarded-For` header. This to be able to deal with a series of reverse
proxies.

For argument's sake, let's assume we have 3 reverse proxies instead of only one:
first one with IP `10.0.1.25` and the second one with `192.168.1.20` (the last
one is not included in the `X-Forwarded-For` header). If the end-user has the IP
address of `2.52.1.39`, then the application would normally receive the header
`X-Forwarded-For: 2.52.1.39, 10.0.1.25, 192.168.1.20` which correctly leads it
to think the end-user has the IP address `2.52.1.39`. But if the end-user passes
their own `X-Forwarded-For` header, for example `X-Forwarded-For: 127.0.0.1`,
then that is included in this list, resulting in the application receiving
`X-Forwarded-For: 127.0.0.1, 2.52.1.39, 10.0.1.25, 192.168.1.20` which leads it
to think the end-user has the IP address `127.0.0.1`.

The reverse proxy for this challenge in particular was a single Nginx, but this
applies to any reverse proxy such as Apache as well, or any combination thereof.
This is a security hole in the application, or rather the architectural design
itself. The application should not be the one doing the IP filtering, but should
instead be dealt with by properly installed firewalls.
