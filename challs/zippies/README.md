# Wrapping Disaster

- Category: `misc`
- Challenge author: **Kalle Fagerberg**

## Description

Santas wrapping machine got bonkers and used all different papers it could find. Can you help unwrap this mess?

## Building prerequisites

```sh
sudo apt update
sudo apt install make p7zip-full gzip zip lz4 xz-utils

# Seem to have forgotten brotli in challenge. Oops.
#sudo apt install brotli
```

## Writeup

The long extension is a givaway; if you know what they mean:

- `.7z`: 7zip compressed archive (<https://www.7-zip.org/>)
- `.lz4`: LZ4 compression (<https://www.lz4.org/>)
- `.zip`: Well known ZIP compressed archive (<https://en.wikipedia.org/wiki/ZIP_(file_format)>), used and supported near to everywhere.
- `.br`: Brotli compression (<https://github.com/google/brotli>), most commonly used by websites as a GZIP alternative for even smaller downloads.
- `.gz`: GZIP compression (<https://www.gzip.org/>), commonly used as de-facto compression in GNU/Linux eco-system and by websites.
- `.b64`: Non-standard extension of Base-64 encoding. Actually only increases the file size by ~33%, but uses a very limited set of ASCII characters making it very portable.
- `.xz`: XZ compression (<https://tukaani.org/xz/>) which trades off speed for better compression compared to many other algorithms
- `.tar`: TAR archive, sans-compression (<https://en.wikipedia.org/wiki/Tar_(computing)>), used to bundle/archive files and directories, where resulting "tarball" file size is always nearly equal to the sum of its content.


To obtain the `flag.txt` hidden in these nested formats you need the above
dependencies. If on Windows, suggest running inside a Docker container, such as

```sh
docker run --rm -it -v ./:/root/ -w /root ubuntu:20.04 bash
```

```sh
7z x flag.tar.xz.b64.gz.br.zip.lz4.7z

lz4 flag.tar.xz.b64.gz.br.zip.lz4

unzip flag.tar.xz.b64.gz.br.zip

# Seem to have forgotten brotli in challenge. Oops.
#brotli flag.tar.xz.b64.gz.br

gunzip flag.tar.xz.b64.gz

base64 -d flag.tar.xz.b64 > flag.tar.xz

# The -J flag tells tar to use xz
tar -xJvf flag.tar.xz

cat flag.txt
```
