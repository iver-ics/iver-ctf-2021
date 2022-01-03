# 3 laws of Xmas

- Category: `web`
- Challenge author: **Pontus Norrström**

## Description

You need to go to the source to get your head around this search.....
too obvious? :-)

### Connection info

<http://2021.santahack.xyz:42201>

## Writeup

1. Check HTML source
2. Check HTTP headers
3. Check `robots.txt`

```console
$ curl http://2021.santahack.xyz:42201/robots.txt
User-agent: *
Disallow: /cool_project.html

$ curl http://2021.santahack.xyz:42201/cool_project.html
<h2>Alright, alright, you found my hideout... Here's the third part of the flag: D_Int3R3st1Ng!}</h2>
```
