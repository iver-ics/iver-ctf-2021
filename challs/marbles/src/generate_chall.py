#!/usr/bin/python3

import os


def xor_bytes(bytes1, bytes2):
    return bytes(a ^ b for a, b in zip(bytes1, bytes2))


flag = b'haxmas{pesky_little_marbles}'
assert(len(flag) % 4 == 0)

value = b'\x00'*8
values_text = ''

while len(flag) > 0:
    for _ in range(1337):  # sen 1337
        xor = os.urandom(8)
        value = xor_bytes(value, xor)
        values_text += '0x' + xor.hex() + 'ULL,\n'
    xor = xor_bytes(value, b'\x13\x37\x13\x37' + flag[:4])
    flag = flag[4:]
    value = xor_bytes(value, xor)
    # print(f"{value=}")
    values_text += '0x' + xor.hex() + 'ULL,\n'

# print(values_text)

with open('marbles_template.c', 'r') as f:
    program_text = f.read()

program_text = program_text.replace("// REPLACETHIS", values_text)

with open('marbles.c', 'w') as f:
    f.write(program_text)
