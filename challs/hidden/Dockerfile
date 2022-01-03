FROM ubuntu:16.04

RUN apt-get update -y && apt-get install -y \
    lib32z1 xinetd

RUN useradd -m ctf

WORKDIR /home/ctf

RUN mkdir /home/ctf/root && \
    mkdir /home/ctf/root/bin && \
    cp /bin/sh /home/ctf/root/bin && \
    cp /bin/ls /home/ctf/root/bin && \
    cp /bin/cat /home/ctf/root/bin

RUN cp -R /lib* /home/ctf/root && \
    cp -R /usr/lib* /home/ctf/root

COPY ./ctf.xinetd /etc/xinetd.d/ctf
COPY ./start.sh /start.sh
RUN echo "Blocked by ctf_xinetd" > /etc/banner_fail

RUN chmod +x /start.sh

RUN mkdir /home/ctf/root/dev && \
    mknod /home/ctf/root/dev/null c 1 3 && \
    mknod /home/ctf/root/dev/zero c 1 5 && \
    mknod /home/ctf/root/dev/random c 1 8 && \
    mknod /home/ctf/root/dev/urandom c 1 9 && \
    chmod 666 /home/ctf/root/dev/*

COPY ./src/flag.txt /home/ctf/root/.flag

RUN chown -R root:ctf /home/ctf && \
    chmod -R 750 /home/ctf && \
    chmod 740 /home/ctf/root/.flag

CMD ["/start.sh"]

EXPOSE 9999