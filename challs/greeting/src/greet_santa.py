#!/usr/bin/env python3

import random
import math
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad
import sys
import os

with open(sys.path[0] + "/flag.txt", "rb") as f:
    flag = f.read().strip()

iv, key = [os.urandom(16) for _ in '01']


def encrypt(text):
    cipher = AES.new(key, AES.MODE_CBC, iv=iv)
    enc = cipher.encrypt(pad(text, 16))
    return enc


def hear():
    return sys.stdin.readline().strip()


def tell(*args):
    s = " ".join(map(str, args))
    sys.stdout.write(s + "\n")
    sys.stdout.flush()


def main():

    border = "*"
    tell(border*72)
    tell(border, " Hi! I want to give you a nice flag as a christmas gift    ")
    tell(border, " I'm gonna wrap it in this super nice AES-paper.           ")
    tell(border*72)

    while True:
        tell("What would you like me to write as a greeting?")
        msg_inp = hear()
        enc = encrypt(str.encode(msg_inp) + flag).hex()
        tell("Awesome! Here's your wrapped flag:\n" + enc)
        tell("\nOh, so you want another gift!")


if __name__ == '__main__':
    main()
