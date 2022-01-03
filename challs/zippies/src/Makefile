
flag:
	tar -cJf flag.tar.xz flag.txt
	base64 -w0 < flag.tar.xz > flag.tar.xz.b64
	gzip -fk flag.tar.xz.b64
	brotli -fk flag.tar.xz.b64.gz
	rm -f flag.tar.xz.b64.gz.br.zip
	zip flag.tar.xz.b64.gz.br.zip flag.tar.xz.b64.gz.br
	lz4 -f -1 flag.tar.xz.b64.gz.br.zip flag.tar.xz.b64.gz.br.zip.lz4
	rm -f flag.tar.xz.b64.gz.br.zip.lz4.7z
	7z a flag.tar.xz.b64.gz.br.zip.lz4.7z flag.tar.xz.b64.gz.br.zip.lz4
