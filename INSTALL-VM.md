# CTFd

## OS

Ubuntu Server 21.10 - Gen2

```sh
sudo apt update -y && sudo apt upgrade -y
```

## CTFd install

```sh
sudo apt install python3-pip -y


git clone https://github.com/bjornosterman/iver_ctf_2021.git

sudo apt update -y && sudo apt install docker.io docker-compose net-tools ncat

pip install ctfcli


sudo docker build -t print_santa .
sudo docker run -d -p "0.0.0.0:42001:9999" --name print_santa print_santa

sudo docker rm -f print_santa
```
