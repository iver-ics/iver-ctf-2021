# Santa's Broke

- Category: `web`
- Challenge author: **Kalle Fagerberg**

## Description

Can you help Santa out of the finacial gopher hole?

### Connection info

```sh
nc 2021.santahack.xyz 70
```

## Writeup

There are different ways to implement this. Either you could create a simple
Gopher protocol parser that finds the equation of each Gopher page, evaluates
it, and then continues. Or, as we'll discuss here is a brute-force
implementation instead.

The Gopher website (or more precisely named: "Gopher hole") has a flaw in that
if you navigate to an incorrect answer then you're instantly presented with the
fact that it was wrong. And as there's only 1 answer that all math equations
lead to, we can write a simple [crawler](https://en.wikipedia.org/wiki/Web_crawler)
that just attempts all links it finds until it finds the flag.

For example: (file can also be found in [`./writeup/solve.py`](./writeup/solve.py))

```python
#!/usr/bin/env python3
from pwn import *
from base64 import b64decode
import re

context.log_level = "ERROR"


def get(selector):
    with remote(b'2021.santahack.xyz', 70) as t:
        t.sendline(selector.encode())
        text = t.recvallS()
        if 'flag' in selector:
            match = re.search('TRAN ID:  \\b(?P<b64>.*?)\\s', text)
            flag = b64decode(match.group('b64'))
            print()
            print(flag.decode())
            exit()
        return re.findall("/(?:flag|math)/[a-z0-9-]+", text)


def next(selectors):
    for selector in selectors:
        next_selectors = get(selector)
        if next_selectors:
            print('.', end='')
            return next_selectors


selector = ['/']

while 1:
    selector = next(selector)
```

The above solution makes use of the `pwntools` framework: <https://docs.pwntools.com/en/stable/>

Can be installed via pip: (docs: <https://docs.pwntools.com/en/stable/install.html#python3>)

```sh
python3 -m pip install --upgrade pwntools
```

Example execution of the script:

```console
$ python3 writeup/solve.py
................................................................................................................................
haxmas{gemini1965-is-much-better-anyway}
```
