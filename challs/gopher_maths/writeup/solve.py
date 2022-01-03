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
