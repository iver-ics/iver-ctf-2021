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
