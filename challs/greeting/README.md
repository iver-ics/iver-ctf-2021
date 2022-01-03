# Greetings

- Category: `crypto`
- Challenge author: **Björn Österman**

## Description

Do you want it wrapped?

### Connection info

```sh
nc 2021.santahack.xyz 42003
```

## Writeup

This challenge can be solved with the insight that the IV and Key is reused for every message.
By starting out by almost filling up a 16-byte segment (ie with 15 chars) and letting the oracle (ie the script running on the server) adding the 16:th character you can then try the same thing for every printable character until you find the same crypto-bytes.
Then iterating one character at the time until you get the whole flag.

```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from pwn import *

io = remote("2021.santahack.xyz", 42003)

buflen = 48
flag = b''


def split_print(msg, end='\n'):
    print(msg[:16] + ' ' + msg[16:32] + ' ' + msg[32:], end=end)


def encrypt(x):
    io.sendlineafter(b"?", x)
    io.recvuntil(b"flag:\n")
    return io.recvline()[:buflen*2]


while not flag[-1:] == b'}':
    text = b'.' * (buflen - len(flag) - 1)
    correct = encrypt(text)
    msg = (text + flag + b'?').decode()
    split_print(msg + ' -> ' + correct[64:72].hex() + '... - Searching...')
    split_print(msg[:47] + ' '*24, '')

    for c in [bytes([ord(x)]) for x in string.printable]:
        check = encrypt(text + flag + c)

        print('\b'*24, end='')
        print(c.decode() + ' -> ' + check[64:72].hex() + '...', end='')

        if check == correct:
            flag += c
            print(" - Found!\n")
            break

print('\n')
```

## Existing writeups

- <https://github.com/EiNSTeiN-/chosen-plaintext>, recommended by Zephyro Kemstedt
- <https://derekwill.com/2021/01/01/aes-cbc-mode-chosen-plaintext-attack/>, recommended by Zephyro Kemstedt
- <https://github.com/ipv6-feet-under/WriteUps-S.H.E.L.L.CTF21/tree/main/Cryptography/Vuln-AES>, recommended by Erik Einarsson
- <https://godiego.co/posts/AES-128-padding-attack/>, recommended by Johan Sjöberg
