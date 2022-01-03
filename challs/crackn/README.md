# Crack Open

- Category: `crypto`
- Challenge author: **Björn Österman**

## Description

That present's gonna rock you!

## Writeup

This password-protected zip-file can be cracked by using the rockyou-wordlist and your favorite hash-cracking utility. In my case 'hashcat'

```sh
/usr/sbin/zip2john crackn.zip > crackn.hash
hashcat --help | grep -i 'pkzip'
hashcat -a 0 -m 17210 --username crackn.hash rockyou.txt
```
