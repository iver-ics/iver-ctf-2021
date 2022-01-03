# Lost his marbles

- Category: `rev`
- Challenge author: **Björn Österman**

## Description

Santa lost some leet marbles in here

## Building prerequisites

```sh
sudo apt update
sudo apt install make clang
```

## Writeup

The challenge is a program where lots of values that are xor:ed together once in a white (actually every 1337:th time) converges into a prefix of hex 0x13371337 and suffix consisting of 4 charactars from the flag.

Every time the prefix is found it prints a "Whoop" and nothing more. But as an unused argument to the printf-function there's the piece of the flag.

To solve this you can start the application in the "gdb"-debugger and add a breakpoint right before it executes the printf call. When the breakpoint hits, the piece of the flag is right there in the RDI-registry.

Keep on "continue" and pick up the pieces build them into the flag.
