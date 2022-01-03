# Lost packet

- Category: `forensics`
- Challenge author: **Robert Teir**

## Description

The hacker that hacked santa downloaded two popular hacking tools, what are the
filenames? `haxmas{filename1.ext_filename2.ext}`

## Writeup

1. Download zip

2. Load pcap in wireshark

3. Load sslkeys in wireshark

4. Find the HTTP GET requests for mimikatz and lazange, resulting in either

   - `haxmas{minikatz.zip_lazange.exe}`
   - _or_ `haxmas{lazange.exe_minikatz.zip}`
