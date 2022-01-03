# Notes

This repo has been a great inspiration for making the setup and automations.
<https://github.com/csivitu/ctf-challenges>

To install `ctfcli` <https://github.com/CTFd/ctfcli>

```sh
pip install ctfcli
```

The template in this repo has been used to host the challenges inside xinetd inside docker 
<https://github.com/Eadom/ctf_xinetd>

CTFd is platform we used for hosting the CTF-web
<https://github.com/CTFd/CTFd>


During the CTF we recorded all packets to and from the challenges with TCPDump. The following 2 shell-scripts where used:
```
[start_tcpdump.sh]
sudo tcpdump -w ~/dumps/challs.pcap -C 10 -Z bjorn 'portrange 42001-42022 or portrange 42201-42204' & disown

[stop_tcpdump.sh]
sudo kill $(pidof tcpdump)
```

## Update IP whitelist

1. Goto: <https://2021.santahack.xyz/admin/config>

2. Click "Backup"

3. Click "Export"

4. Extract the file "tracking.json"

5. Run the following:

   ```sh
   echo $(jq ".results[].ip" tracking.json -r | sort | uniq ) | tr ' ' ',' | clip.exe
   ```

6. Goto Azure Portal and lookup the NSG for the Challenge-server

7. Goto: "Inbound Security Rules"

8. Add a new rule

   - From IP Addresses: [paste from clipboard here]
   - To IP Addresses: All
   - To ports: 70,42001-42299
   - Protocol: TCP

We added a rule each time, but could have replace the rule instead.

Next time we will probably automate the process with a cron-job.

