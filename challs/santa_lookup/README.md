# Santa Policy Framework

- Category: `OSINT`
- Challenge author: **Robert Teir**

## Description

Santa has hidden the flag where everyone can look it up.

## Writeup

Found as a TXT DNS record. For making it easier, the following domains all have
the flag as a TXT record:

- 2021.santahack.xyz
- flag.santahack.xyz
- santahack.xyz

Could be done via free online service, such as
<https://mxtoolbox.com/SuperTool.aspx?action=txt%3a2021.santahack.xyz&run=toolpage>,
or via the command-line, such as via `nslookup` or `dig`:

```console
$ nslookup -query=TXT 2021.santahack.xyz
Server:		127.0.0.53
Address:	127.0.0.53#53

Non-authoritative answer:
2021.santahack.xyz	text = "haxmas{so_you_have_heard_about_txt_records}"

$ dig santahack.xyz -t TXT +noall +answer
2021.santahack.xyz.	600	IN	TXT	"haxmas{so_you_have_heard_about_txt_records}"
```

The challenge name, "Santa Policy Framework", is a reference to SPF, which is
a TXT DNS record format used in email communication and verification.
