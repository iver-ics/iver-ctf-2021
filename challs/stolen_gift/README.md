# Stolen Gift

- Category: `misc`
- Challenge author: **Pontus Norrström**

## Description

Santa just turned 1e37 years old and has a bit of a mid life crisis and bought
a hoj, only to have it stolen! Miscreants! Fortunately he had a tracker
attached. Can you help extract the data?

## Writeup

```sh
$ binwalk -e stolen_gift.jpg

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
78184         0x13168         Zip archive data, at least v1.0 to extract, compressed size: 23, uncompressed size: 23, name: flag.txt
78335         0x131FF         End of Zip archive, footer length: 22

$ cat _stolen_gift.jpg.extracted/flag.txt
xmas{n3V3r_tRu5t_f1L3S}
```


or, more crudely you can use **strings**

```sh
$ strings -n 12 stolen_gift.jpg
!22222222222222222222222222222222222222222222222222
%&'()*456789:CDEFGHIJSTUVWXYZcdefghijstuvwxyz
&'()*56789:CDEFGHIJSTUVWXYZcdefghijstuvwxyz
flag.txtxmas{n3V3r_tRu5t_f1L3S}PK
```
The reason this works is because the flag.txt-file that was inside the zip-file is so small that zip uses uncompressed format.

This time is also cuts away the first to character making the flag i little messed up. :-)