# Hacker zipping glögg

- Category: `OSINT`
- Challenge author: **Pontus Norrström**

## Description

Santas Offensive Security Team has found a picture of the aformidable
adversary. Can you help them find out where he was located at the time?
Flag format: `haxmas{The_Location}`

## Writeup

1. Check location metadata on image, ex via `exiftool`
2. Look up the location in e.g Google Maps
3. Flag name is the name of the cafe, i.e. "Cyber Link" => `haxmas{cyber_link}`

Example extraction of location metadata:

```sh
exiftool -c '%.6f' -GPSPosition image.jpg
```

The most difficult part with this challenge was actually the fact that Google Maps hides the name of the cafe to instead show the marker from the GPS-coordinates. Removing the search-sidebar after finding the location reveals the Cyber Link pin.